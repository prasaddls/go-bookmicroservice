mkdir go-bookmicroservice
cd go-bookmicroservice
go mod init github.com/prasaddls/go-bookmicroservice
go get "github.com/jinzhu/gorm"
go get "github.com/jinzhu/gorm/dialects/mysql"
go get "github.com/gorilla/mux"


git remote add origin git@github.com:prasaddls/go-bookmicroservice.git
git branch -M main
git add .
git commit -m "initial commit"
git push -u origin main


go mod tidy  to download all packages 