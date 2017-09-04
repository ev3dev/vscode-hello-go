# vscode-hello-go
Hello World for [ev3dev] + [Visual Studio Code][code] + [Go]


## Overview
This is a git repository to help you get started programming a robot using
ev3dev in Visual Studio Code using the Go programming language.


## Prerequisites
* LEGO MINDSTORMS EV3, Dexter Industries Brick Pi, Mindsensors PiStorms, or
  FatcatLab EVB with [ev3dev] installed. **Does not work with ev3dev-jessie!
  Be sure to grab a snapshot image of ev3dev-stretch.**
* A computer (Windows, macOS, or Linux) with [Go], [Visual Studio Code][code] and
  [git] (or [GitHub Desktop][github]) installed.


## Step-by-Step
1.  Start Visual Studio Code.

2.  Select *Clone Git repository...* in the *Welcome* tab.

    ![screenshot](.README/vscode-welcome-clone-git.png)
3.  Enter `https://github.com/ev3dev/vscode-hello-go.git` as the repository
    name and press <kbd>Enter</kbd>.

    ![screenshot](.README/vscode-repository-url.png)

4.  Enter the name of the directory where you would like to clone (download)
    this repository. This can be any place you like to save files.

    ![screenshot](.README/vscode-parent-directory.png)

5. Click *Open Repository* when asked.

    ![screenshot](.README/vscode-open-repository.png)

6.  Click *Show Recommendations* when asked.

    ![screenshot](.README/vscode-show-recommendations.png)

7.  Install the `ev3dev-browser` extension. The Go extension is optional, but it
    is nice to have too.

    ![screenshot](.README/vscode-extensions-recommended.png)

8.  After installation completes, click *Reload* and *Reload Window*.

    ![screenshot](.README/vscode-ev3dev-browser-reload.png)

    ![screenshot](.README/vscode-reload-window.png)

9. Open the *Explorer* activity pane.

    ![screenshot](.README/vscode-explorer-icon.png)

10. Click the arrow next to *EV3DEV DEVICE BROWSER* to open it.

    ![screenshot](.README/vscode-ev3dev-browser-collapsed.png)

    ![screenshot](.README/vscode-ev3dev-browser-expanded.png)

11. Ensure that your ev3dev device is turned on and has a network connection to
    the host computer. USB, Bluetooth, Wi-Fi or wired will work.

12. Click the text where it says "Click here to connect to a device".
    A box will pop up that lists discovered devices. Select one.

    ![screenshot](.README/vscode-searching-for-devices.png)

13. Once the device has connected, you will see a green dot and the
    `/home/robot` folder.

    ![screenshot](.README/vscode-home-robot.png)

14. Press <kbd>F5</kbd> to download the program and run it.

15. The *Output* pane will automatically open and show the status of the
    program and any error/debug messages.

    ![screenshot](.README/vscode-output.png)

    ![screenshot](.README/ev3-output.png)

18. Open the source code file, [hello.go](./hello.go), to learn more.


[ev3dev]: http://www.ev3dev.org
[code]: https://code.visualstudio.com/
[go]: https://golang.org/
[git]: https://git-scm.com/
[github]: https://desktop.github.com/
