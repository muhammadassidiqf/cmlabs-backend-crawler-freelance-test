# Introduction
Ini merupakan RESTful API untuk Crawler Website.\
Pertama clone aplikasi ini menggunakan command dibawah ini.
```bash
# Download this project
git clone https://github.com/muhammadassidiqf/cmlabs-backend-crawler-freelance-test.git
```

# Build & Run
```bash
cd cmlabs-backend-crawler-freelance-test
# Build
docker build -t crawler-service:latest -f Dockerfile .
# Run
docker run --rm --name crawler-service -v "C:\Data\crawler-website:/app" -p 7777:7777 crawler-service
# API Endpoint : http://localhost:7777
```

# API

### http://localhost:7777/cmlabs
* `GET` : Get Website cmlabs

### http://localhost:7777/sequence
* `GET` : Get Website sequence

### http://localhost:7777/posibel
* `GET` : Get Website posibel
