# quick-dict
Quickly translate the selected words.

## Requiements
This tool uses [translate-shell](https://github.com/soimort/translate-shell) to translate text. So you must install trans first.
```sh
sudo apt install translate-shell
```

## Installation



Run 
```sh
go install github.com/thuan1412/quick-dict@latest
```
to install the executable file.

Or build from source code.

In order to listen to the speak of the text, install [google-speech](https://pypi.org/project/google-speech/)
and its dependencies

```sh
pip install google-speech
sudo apt-get install sox
sudo apt-get install libsox-fmt-mp3
```

And create a shortcut to run the command.

![image](https://user-images.githubusercontent.com/36019052/233822979-fe205d12-d59e-463d-896c-1c47bcbaaec5.png)


Then, you can use the shortcut to open the popup window to translate the selected text to Vietnamese.



## Demo
After install successfully, you can use the shortcut the translate the selected text
![image](https://user-images.githubusercontent.com/36019052/234150381-a6ba19ac-451b-4406-9c91-c75ca2fbf48f.png)
