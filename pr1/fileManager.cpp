#include "fileManager.h"

auto FileManager::openFile() -> std::optional<std::tuple<QTextEdit *, QString>> {
    const auto fileName = QFileDialog::getOpenFileName(this,
                                                       "Open file",
                                                       QDir::currentPath());
    if (fileName.isEmpty()) {
        return std::nullopt;
    }

    for (const auto& i : fileList_.values()) {
        if (i->fileName() == fileName) {
            return std::nullopt;
        }
    }

    const auto file = new QFile(fileName);
    file->open(QIODevice::ReadOnly);
    const auto data = QString::fromUtf8(file->readAll());
    file->close();

    const auto textEdit = new QTextEdit();
    textEdit->setPlainText(data);

    fileList_.insert(textEdit, file);

    return std::tie(textEdit, fileName);
}

bool FileManager::closeFile(QTextEdit *textEdit) {
    bool fileChanged = false;
    bool fileIsNew = !fileList_.contains(textEdit);

    if (!fileIsNew) {
        const auto file = fileList_[textEdit];
        file->open(QIODevice::ReadOnly);
        const auto data = file->readAll();
        file->close();

        if (data != textEdit->toPlainText()) {
            fileChanged = true;
        }
    }

    if (!fileIsNew && !fileChanged) {
        fileList_.remove(textEdit);
        return true;
    }

    if (fileIsNew || fileChanged) {
        QMessageBox::StandardButton reply;
        reply = QMessageBox::question(this, "Save file", "Save file?",
                                      QMessageBox::Yes | QMessageBox::No | QMessageBox::Cancel, QMessageBox::Cancel);
        if (reply == QMessageBox::Yes || reply == QMessageBox::No) {
            if (reply == QMessageBox::Yes) {
                if (!saveFile(textEdit)) {
                    return false;
                }
            }

            fileList_.remove(textEdit);
            return true;
        }
    }

    return false;
}

std::optional<QString> FileManager::saveFile(QTextEdit *textEdit) {
    QFile *file;
    QString fileName;

    if (fileList_.contains(textEdit)) {
        file = fileList_[textEdit];
        fileName = file->fileName();
    } else {
        fileName = QFileDialog::getSaveFileName(this,
                                               QString::fromUtf8("Сохранить файл"),
                                               QDir::currentPath(),
                                               "All files (*.*)");
        if (fileName.isEmpty()) {
            return std::nullopt;
        }

        file = new QFile(fileName);
        fileList_.insert(textEdit, file);
    }

    const auto &data = textEdit->toPlainText();

    if (!file->open(QIODevice::WriteOnly)) {
        return std::nullopt;
    }

    if (data.isEmpty()) {
        file->close();
        return fileName;
    }

    if (!file->write(data.toUtf8())) {
        file->close();
        return std::nullopt;
    }

    file->close();

    return fileName;
}

void FileManager::saveAllFiles() {
    for (const auto textEdit : fileList_.keys()) {
        saveFile(textEdit);
    }
}

