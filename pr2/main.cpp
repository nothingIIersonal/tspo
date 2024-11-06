// Создание службы мониторинга системных ресурсов

// Описание: разработка демона Linux, который собирает информацию о системных
//           ресурсах (использование CPU,
//           память, диск, сеть) и сохраняет данные в лог-файл или базу данных
//           для последующего анализа.
// Цели:     изучение разработки демонов Linux, работа с системным API.

// Реализуемые задачи:

// [+]  • Разработать демон Linux, который автоматически запускается при старте
//        системы.
// [+]  • Реализовать сбор данных об использовании процессора, памяти, дискового
//        пространства и сети.
// [+]  • Добавить возможность записи данных в лог-файл или базу данных с
//        заданной периодичностью.
// [+]  • Реализовать функционал оповещения (например, через email) при
//        превышении заданных пороговых значений использования ресурсов.
// [+]  • Обеспечить возможность конфигурирования службы через файл настроек или
//        интерфейс администратора.
// [+]  • Реализовать систему уведомлений в реальном времени о состоянии службы
//        (начало работы, ошибки и т.д.).
// [+]  • Добавить возможность остановки и перезапуска службы.
// [-]  • Реализовать защиту службы от несанкционированного доступа и изменения
//        настроек.
// [-]  • Создать интерфейс для просмотра и анализа собранных данных.

#include <chrono>
#include <cinttypes>
#include <csignal>
#include <cstring>
#include <ctime>
#include <fstream>
#include <iostream>
#include <thread>

#include "easy-email-cpp/include/easy-email.hpp"
#include "json.hpp"
#include "linux-system-usage.hpp"

bool running = true;

struct settings_t {
  std::string log_to;
  uint64_t update_time_s;
  std::string email;
  std::string net_iface;
  struct threasholds_t {
    uint64_t CPU;
    uint64_t RAM;
    uint64_t memory;
    uint64_t network;
  } threasholds;
} settings;

nlohmann::json out_json = {};
std::ofstream out_file;

// you may need to change smtp server in easy-email-cpp/include/easy-email.hpp
std::string email(enter email here);
std::string pass(enter pass here);
EasyEmail easy_email(email, pass);

void print_threasholds(settings_t::threasholds_t thr) {
  std::cout << "\tCPU: " << thr.CPU << std::endl;
  std::cout << "\tRAM: " << thr.RAM << std::endl;
  std::cout << "\tMemory: " << thr.memory << std::endl;
  std::cout << "\tNetwork: " << thr.network << std::endl;
}

void print_settings() {
  std::cout << "-------------------------------------------\n";
  std::cout << "Loaded data:\n";
  std::cout << "-------------------------------------------\n";
  std::cout << "Logging to file: " << settings.log_to << "\n";
  std::cout << "Update log: " << settings.update_time_s << "\n";
  std::cout << "Notification email: " << settings.email << "\n";
  std::cout << "Network interface: " << settings.net_iface << "\n";
  std::cout << "Threasholds:\n";
  std::cout << "-------------------------------------------\n";
  print_threasholds(settings.threasholds);
}

void log_thr(settings_t::threasholds_t thr) {
  std::time_t now = std::time(nullptr);
  std::tm *localTime = std::localtime(&now);
  std::ostringstream oss;
  oss << std::put_time(localTime, "%Y-%m-%d %H:%M:%S");
  const auto cur = oss.str();

  out_json[cur]["CPU"] = thr.CPU;
  out_json[cur]["RAM"] = thr.RAM;
  out_json[cur]["Memory"] = thr.memory;
  out_json[cur]["Network"] = thr.network;

  static bool first_print = true;

  if (first_print) {
    out_file << "\"" << cur << "\": ";
    first_print = false;
  } else {
    out_file << "," << std::endl << "\"" << cur << "\": ";
  }
  out_file << std::setw(4) << out_json[cur];
}

void load_data(char *path) {
  auto data = nlohmann::json::parse(std::ifstream(path));

  settings.log_to = data["LogTo"];
  settings.update_time_s = data["UpdateTimeS"];
  settings.email = data["Email"];
  settings.net_iface = data["NetIface"];
  settings.threasholds.CPU = data["Threasholds"]["CPU"];
  settings.threasholds.RAM = data["Threasholds"]["RAM"];
  settings.threasholds.memory = data["Threasholds"]["Memory"];
  settings.threasholds.network = data["Threasholds"]["Network"];
}

void check_threasholds(settings_t::threasholds_t entry) {
  std::string msg = "";

  if (entry.CPU > settings.threasholds.CPU) {
    msg += "- Threashold for CPU reached! Value: " + std::to_string(entry.CPU) +
           ", max: " + std::to_string(settings.threasholds.CPU) + "\n";
  }
  if (entry.RAM > settings.threasholds.RAM) {
    msg += "- Threashold for RAM reached! Value: " + std::to_string(entry.RAM) +
           ", max: " + std::to_string(settings.threasholds.RAM) + "\n";
  }
  if (entry.memory > settings.threasholds.memory) {
    msg += "- Threashold for disk memory reached! Value: " +
           std::to_string(entry.memory) +
           ", max: " + std::to_string(settings.threasholds.memory) + "\n";
  }
  if (entry.network > settings.threasholds.network) {
    msg += "- Threashold for network reached! Value: " +
           std::to_string(entry.network) +
           ", max: " + std::to_string(settings.threasholds.network) + "\n";
  }

  if (msg != "") {
    std::cout << "ONE OR MORE THREASHOLDS WERE REACHED!\n" << msg << std::endl;
    easy_email.send(settings.email, "Threasholds", msg);
  }
}

int monitor() {
  out_file.open(settings.log_to);
  if (!out_file.is_open()) {
    std::cerr << "Can't open log file" << std::endl;
    return EXIT_FAILURE;
  }

  out_file << "{" << std::endl;

  settings_t::threasholds_t entry;

  std::cout << "[STARTED DAEMON]" << std::endl;

  while (running) {
    // CPU usage
    const auto t1 = get_system_usage_linux::read_cpu_data();
    std::this_thread::sleep_for(std::chrono::milliseconds(10));
    const auto t2 = get_system_usage_linux::read_cpu_data();
    entry.CPU = 100 * get_system_usage_linux::get_cpu_usage(t1, t2);

    // Memory usage
    entry.RAM =
        100 * get_system_usage_linux::read_memory_data().get_memory_usage();

    // Disk usage
    entry.memory = 100 * get_system_usage_linux::get_disk_usage("/");

    // Network usage
    entry.network =
        get_system_usage_linux::get_network_usage(settings.net_iface);

    // print_threasholds(entry);
    log_thr(entry);

    check_threasholds(entry);

    std::this_thread::sleep_for(
        std::chrono::milliseconds(settings.update_time_s * 1000));
  }

  out_file << "}" << std::endl;

  std::cout << "[STOPPED DAEMON]" << std::endl;

  return 0;
}

void sigterm_handler(int signum) { running = false; }

int main(int argc, char *argv[]) {
  std::cout << "[STARTING DAEMON]" << std::endl;

  if (argc != 2) {
    std::cout << "Set settings JSON path as second argument" << std::endl;
    exit(EXIT_SUCCESS);
  }

  std::signal(SIGTERM, sigterm_handler);

  try {
    load_data(argv[1]);
    print_settings();
  } catch (...) {
    std::cerr << "Can't parse settings" << std::endl;
    exit(EXIT_FAILURE);
  }

  return monitor();
}