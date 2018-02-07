# imgapi
Go Developer - ImgAPI Challenge

1. Prerequisites.

You need to have Docker and Go installed on your machine. Links to download:  
[Download Docker](https://docs.docker.com/install/)  
[Download Go](https://golang.org/dl/)

2. Installation process

Clone the "imgapi" repository and make it:

For Linux and Mac OS run the following commands in Terminal:  
```
    $ go get -u github.com/Kapitoshka438/imgapi  
    $ cd ~/go/src/github.com/Kapitoshka438/imgapi  
    $ make  
```
   
  For Windows OS run the following command in the Command prompt:  
```
    go get -u github.com/Kapitoshka438/imgapi  
    cd %HOMEPATH%\go\src\github.com\Kapitoshka438\imgapi  
    run.bat  
```

3. Usage

Open "localhost:8080" in your web browser.

![ImgAPI Demo Application](screenshot.png?raw=true)

Press the "Choose file" button and choose an image from your machine.  
Press the "Upload image" to upload the chosen file.  
Now you can see the uploaded image in the list of images of the main screen.  
You can Resize, Crop and Rotate images from the list, also using the combinations of these operations and change the format of them.  
You can change the default parameters of these operations in the URL path.
