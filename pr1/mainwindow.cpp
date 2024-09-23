#include "mainwindow.h"
#include "./ui_mainwindow.h"
#include <QPushButton>
#include <QHBoxLayout>
#include <QLabel>
#include <QFileDialog>
#include <QTabWidget>
#include <QMdiSubWindow>

#define ADD_QTEXTEDIT_WINDOW(title, textEdit)                                                  \
                const auto &subWindow = mdiArea_->addSubWindow(textEdit);                      \
                subWindow->setWindowTitle(title);                                              \
                subWindow->setAttribute(Qt::WA_DeleteOnClose);                                 \
                subWindow->setWindowFlags(Qt::WindowTitleHint | Qt::WindowMinimizeButtonHint); \
                subWindow->show()

void MainWindow::newFile_() {
    qDebug() << "newFile";

    ADD_QTEXTEDIT_WINDOW("untitled", new QTextEdit());
}

void MainWindow::openFile_() {
    qDebug() << "openFile";

    const auto ret = fileManager_->openFile();
    if (!ret) {
        msgBox_.critical(this, "Error", "Can't open file. It may be already open.");
        return;
    }

    const auto [textEdit, fileName] = ret.value();
    ADD_QTEXTEDIT_WINDOW(fileName, textEdit);
}

void MainWindow::saveFile_() {
    qDebug() << "saveFile";

    const auto activeSubWindow = mdiArea_->activeSubWindow();
    if (activeSubWindow == nullptr) {
        return;
    }

    const auto widget = activeSubWindow->widget();
    const auto ret = fileManager_->saveFile((QTextEdit *)widget);
    if (!ret) {
        msgBox_.critical(0, "Error", "Can't save file");
        return;
    }

    const auto fileName = ret.value();
    activeSubWindow->setWindowTitle(fileName);
}

void MainWindow::closeFile_() {
    qDebug() << "closeFile";

    const auto activeSubWindow = mdiArea_->activeSubWindow();
    if (activeSubWindow == nullptr) {
        return;
    }

    const auto widget = activeSubWindow->widget();
    if (fileManager_->closeFile((QTextEdit *)widget)) {
        mdiArea_->closeActiveSubWindow();
    }
}

void MainWindow::autosave_() {
    qDebug() << "autoSave";

    if (ui_->actionAutosave->isChecked()) {
        timer_->start(5000);
    } else {
        timer_->stop();
    }
}

void MainWindow::saveAllFiles_() {
    qDebug() << "saveAllFiles";

    fileManager_->saveAllFiles();
}

void MainWindow::printFile_() {
    qDebug() << "printFile_";

    const auto activeSubWindow = mdiArea_->activeSubWindow();
    if (activeSubWindow == nullptr) {
        return;
    }

    dialog_->setWindowTitle("Print " + activeSubWindow->windowTitle());

    if (dialog_->exec() != QDialog::Accepted) {
        msgBox_.critical(0, "Error", "Can't print file");
        return;
    }

    const auto textEdit = (QTextEdit *)activeSubWindow->widget();

    painter_.begin(&printer_);
    painter_.drawText(100, 100, 500, 500, Qt::AlignLeft|Qt::AlignTop, textEdit->toPlainText());
    painter_.end();
}

void MainWindow::findText_() {
    qDebug() << "findText_";

    const auto activeSubWindow = mdiArea_->activeSubWindow();
    if (activeSubWindow == nullptr) {
        return;
    }

    const auto textEdit = (QTextEdit *)activeSubWindow->widget();

    bool ok = false;
    QString toFind = QInputDialog::getText(this, tr("Find text"), tr("Enter text to find"), QLineEdit::Normal, 0, &ok);
    if (ok && !toFind.isEmpty()) {
        textEdit->find(toFind);
    }
}

void MainWindow::findAndReplaceText_() {
    qDebug() << "findAndReplaceText_";
}

void MainWindow::setTextColor_() {
    qDebug() << "setTextColor";

    const auto activeSubWindow = mdiArea_->activeSubWindow();
    if (activeSubWindow == nullptr) {
        return;
    }

    const auto textEdit = (QTextEdit *)activeSubWindow->widget();

    colorDialog_->setWindowTitle("Select color for " + activeSubWindow->windowTitle());
    QColor color = colorDialog_->getColor();

    if (!color.isValid()) {
        msgBox_.critical(0, "Error", "Can't set text color");
        return;
    }

    textEdit->setTextColor(color);
}

void MainWindow::setupConnects_() {
    connect(ui_->actionNew, SIGNAL(triggered()), this, SLOT(newFile_()));
    connect(ui_->actionSave, SIGNAL(triggered()), this, SLOT(saveFile_()));
    connect(ui_->actionOpen, SIGNAL(triggered()), this, SLOT(openFile_()));
    connect(ui_->actionClose, SIGNAL(triggered()), this, SLOT(closeFile_()));
    connect(ui_->actionAutosave, SIGNAL(triggered()), this, SLOT(autosave_()));
    connect(ui_->actionColor, SIGNAL(triggered()), this, SLOT(setTextColor_()));
    connect(timer_, SIGNAL(timeout()), this, SLOT(saveAllFiles_()));
}

void MainWindow::setupShortcuts_() {
    // Ctrl + N
    keyCtrlN_ = new QShortcut(this);
    keyCtrlN_->setKey(Qt::CTRL | Qt::Key_N);
    connect(keyCtrlN_, SIGNAL(activated()), this, SLOT(newFile_()));

    // Ctrl + S
    keyCtrlS_ = new QShortcut(this);
    keyCtrlS_->setKey(Qt::CTRL | Qt::Key_S);
    connect(keyCtrlS_, SIGNAL(activated()), this, SLOT(saveFile_()));

    // Ctrl + O
    keyCtrlO_ = new QShortcut(this);
    keyCtrlO_->setKey(Qt::CTRL | Qt::Key_O);
    connect(keyCtrlO_, SIGNAL(activated()), this, SLOT(openFile_()));

    // Ctrl + W
    keyCtrlW_ = new QShortcut(this);
    keyCtrlW_->setKey(Qt::CTRL | Qt::Key_W);
    connect(keyCtrlW_, SIGNAL(activated()), this, SLOT(closeFile_()));

    // Ctrl + P
    keyCtrlP_ = new QShortcut(this);
    keyCtrlP_->setKey(Qt::CTRL | Qt::Key_P);
    connect(keyCtrlP_, SIGNAL(activated()), this, SLOT(printFile_()));

    // Ctrl + F
    keyCtrlF_ = new QShortcut(this);
    keyCtrlF_->setKey(Qt::CTRL | Qt::Key_F);
    connect(keyCtrlF_, SIGNAL(activated()), this, SLOT(findText_()));

    // Ctrl + H
    keyCtrlH_ = new QShortcut(this);
    keyCtrlH_->setKey(Qt::CTRL | Qt::Key_H);
    connect(keyCtrlH_, SIGNAL(activated()), this, SLOT(findAndReplaceText_()));

    // Ctrl + L
    keyCtrlL_ = new QShortcut(this);
    keyCtrlL_->setKey(Qt::CTRL | Qt::Key_L);
    connect(keyCtrlL_, SIGNAL(activated()), this, SLOT(setTextColor_()));
}

MainWindow::MainWindow(QWidget *parent)
    : QMainWindow(parent)
    , ui_(new Ui::MainWindow)
    , mdiArea_(new QMdiArea(this))
    , fileManager_(new FileManager)
    , timer_(new QTimer(this))
{
    ui_->setupUi(this);
    ui_->actionAutosave->setCheckable(true);

    dialog_ = new QPrintDialog(&printer_);
    colorDialog_ = new QColorDialog(this);

    setupConnects_();
    setupShortcuts_();

    setCentralWidget(mdiArea_);
}

MainWindow::~MainWindow()
{
    delete ui_;
}
