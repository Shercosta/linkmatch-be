go mod tidy
cd prisma
npm install
npx prisma migrate dev --name auto_first_time
cd ..
swag init
# go run main.go