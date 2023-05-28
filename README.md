# quick-dict
Quickly translate the selected words.

## Features
- Translate the selected text using [translate-shell](https://github.com/soimort/translate-shell)
- Speak the selected text using [google-speech](https://pypi.org/project/google-speech/)

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
![image](https://github.com/thuan1412/quick-dict/assets/36019052/904b5f6d-0404-42c4-acd3-23d5a9aca6e6)
