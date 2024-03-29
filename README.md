### vieclamit - bot tìm việc làm IT
- Nguồn tin tuyển dụng:
    * [x]  [topcv](https://www.topcv.vn/tim-viec-lam-it-phan-mem-c10026)
    * [ ]  [jobsgo](https://jobsgo.vn/viec-lam-cong-nghe-thong-tin.html)
    
- Yêu cầu 1 tin tuyển dụng (recruitment) hợp lệ phải đủ các trường sau:
```
{
    "title" : "Sales Manager (IT Outsourcing) - Up To 2500$",
    "company" : "Công ty CP Savvycom",
    "location" : "Hà Nội",
    "salary" : "Tới 2,500 USD",
    "url_job" : "https://www.topcv.vn/brand/savvycomsoftware/tuyen-dung/sales-manager-it-outsourcing-up-to-2500-j595803.html",
    "url_company" : "https://savvycomsoftware.com/",
    "job_deadline" : "15/03/2022"
} 
```

- Tiêu chí:
* [x]  Thu thập hết dữ liệu từ nguồn
* [x]  Dữ liệu thu thập không bị trùng lặp
* [x]  Lập lịch tự động thu thập dữ liệu
* [x]  Lập lịch tự động xóa các tin tuyển dụng quá hạn

- Chức năng:
    - Tìm kiếm tin tuyển dụng theo từ khóa không phân biệt chữ hoa/thường, phải đủ dấu:
        * [x]  Từ khóa: skill (golang, python, php,...)
        * [x]  Từ khóa: location (Hà nội, Hồ chí minh, đà nẵng,...)
        * [x]  Từ khóa: company (vccorp, FPT, vng,...)
        * [x]  Từ khóa: location (Hà nội, Hồ chí minh, đà nẵng,...) và skill (golang, python, php,...)  

- Cài đặt:
    - [golang-install](https://go.dev/doc/install)
    - [mongodb-on-ubuntu-20-04](https://www.digitalocean.com/community/tutorials/how-to-install-mongodb-on-ubuntu-20-04)
    - [setup-mongodb-atlas-deploy-heroku](https://www.mongodb.com/developer/how-to/use-atlas-on-heroku/)
- Sử dụng:
```
$ go run main.go
```
    Hoặc:
```
$ go build
$ ./vieclamit
```

- Deploy app to heroku
 
```
$ heroku login
$ heroku create vieclamit
$ heroku config:set MONGO_URI=
$ heroku config:set TELEGRAM_TOKEN=
$ heroku config:set DATABASE_NAME=
$ heroku config:set COLLECTION=

$ cd my-project/
$ git init
$ heroku git:remote -a vieclamit
$ heroku stack:set container
$ git status
$ git add .
$ git commit -am "make it better"
$ git push heroku master
$ heroku ps:scale worker=1

$ heroku logs --tail
```
![alt text](https://github.com/dactoankmapydev/vieclamit/blob/master/doc_pictures/vli.png)

- Sử dụng bot **vieclamit** trên telegram:
    - Bắt đầu và hướng dẫn sử dụng:
    
    - Tìm kiếm tin tuyển dụng theo từ khóa công ty:
    
    ![alt text](https://github.com/dactoankmapydev/vieclamit/blob/master/doc_pictures/company.png)
    
    - Tìm kiếm tin tuyển dụng theo từ khóa kỹ năng:
    
    ![alt text](https://github.com/dactoankmapydev/vieclamit/blob/master/doc_pictures/skill.png)
    
    - Tìm kiếm tin tuyển dụng theo từ khóa địa điểm và kỹ năng:
    
    ![alt text](https://github.com/dactoankmapydev/vieclamit/blob/master/doc_pictures/vli2.png)
