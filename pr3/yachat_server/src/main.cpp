// Разработка простого чата с использованием сокетов:

// Описание: проект предусматривает создание серверного и клиентского приложения
// для обмена сообщениями в локальной сети. Сервер должен поддерживать несколько
// клиентов, а клиенты должны иметь возможность отправлять и получать сообщения.
// Цели: изучение межпроцессного взаимодействия и работы с сокетами.

// Реализуемые задачи:
// [+] • Реализовать серверное приложение, способное обрабатывать подключения от
//       нескольких клиентов.
// [*] • Реализовать клиентское приложение с возможностью подключения к серверу
//       по IP-адресу и порту.
// [+] • Обеспечить двусторонний обмен сообщениями между клиентом и сервером.
// [-] • Добавить поддержку отправки и получения файлов через чат.
// [*] • Реализовать возможность выбора отображения чата в виде окна с
//       прокруткой и цветовым выделением сообщений от разных пользователей.
// [+] • Добавить возможность приватных сообщений между клиентами.
// [+] • Реализовать систему авторизации и регистрации пользователей.
// [+] • Добавить журналирование сообщений на стороне сервера для последующего
//       анализа.
// [-] • Обеспечить защиту данных при передаче (например, через шифрование
//       сообщений).

#include <QCoreApplication>
#include <QString>

#define JOURNAL "journal.txt"

#include "server.hpp"

int main(int argc, char *argv[]) {
  QCoreApplication a(argc, argv);

  server::Server server(1234, &a);
  common::logAll(QtDebugMsg, "[MAIN] Server started on port " +
                                 QString::number(server.serverPort()));

  return a.exec();
}
