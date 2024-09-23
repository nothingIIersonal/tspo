#ifndef FILEMANAGER_H
#define FILEMANAGER_H

#include <QWidget>
#include <QTextEdit>
#include <QMap>
#include <QPair>
#include <QFile>
#include <QFileDialog>
#include <QSaveFile>
#include <QMessageBox>

#include <optional>
#include <fstream>
#include <tuple>

class Q_DECL_EXPORT FileManager : public QWidget
{
    Q_OBJECT

public:
    FileManager() = default;
    bool closeFile(QTextEdit *textEdit);
    std::optional<std::tuple<QTextEdit *, QString>> openFile();
    std::optional<QString> saveFile(QTextEdit *textEdit);
    void saveAllFiles();
private:
    QMap<QTextEdit *, QFile *> fileList_;
};

#endif // FILEMANAGER_H
