#ifndef MAINWINDOW_H
#define MAINWINDOW_H

#include <QMainWindow>
#include <QShortcut>
#include <QMdiArea>
#include <QTimer>
#include <QColorDialog>
#include <QLineEdit>
#include <QInputDialog>
#include <QPainter>
#include <QtPrintSupport/QPrinter>
#include <QtPrintSupport/QPrintDialog>

#include "fileManager.h"

QT_BEGIN_NAMESPACE
namespace Ui {
class MainWindow;
}
QT_END_NAMESPACE

class MainWindow : public QMainWindow
{
    Q_OBJECT

public:
    MainWindow(QWidget *parent = nullptr);
    ~MainWindow();
private slots:
    void newFile_();
    void saveFile_();
    void openFile_();
    void closeFile_();
    void autosave_();
    void saveAllFiles_();
    void printFile_();
    void findText_();
    void findAndReplaceText_();
    void setTextColor_();
private:
    void setupConnects_();
    void setupShortcuts_();
    Ui::MainWindow *ui_;
    QMdiArea *mdiArea_;
    FileManager *fileManager_;
    QTimer *timer_;
    QPrintDialog *dialog_;
    QShortcut *keyCtrlN_;
    QShortcut *keyCtrlS_;
    QShortcut *keyCtrlO_;
    QShortcut *keyCtrlW_;
    QShortcut *keyCtrlP_;
    QShortcut *keyCtrlF_;
    QShortcut *keyCtrlH_;
    QShortcut *keyCtrlL_;
    QMessageBox msgBox_;
    QPrinter printer_;
    QPainter painter_;
    QColorDialog *colorDialog_;
};

#endif // MAINWINDOW_H
