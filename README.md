# go-word-references
An API REST service to help you to improve your english vocabulary

# Why in Go?
Because I was curious about try this framework: https://github.com/gin-gonic/gin

# How can i run it?
There're are three ways for run this project:

First you have to go to the project folder

1. Simple run `go run main.go`
2. By using https://github.com/codegangsta/gin it will create a proxy server that will listen the changes of your code. With this option you have to run `gin -i`
3. Run `docker-compos up`with a `.env`file previusly created.  **If you're using this option, just remember the variables name, you can check it in the docker-compose.yml file, they're the ones with the `${}` characters**
4. Run it with Air Live reload: 
