##
## STEP 1 - BUILD
##

# specify the base image to  be used for the application, alpine or ubuntu
FROM golang:1.21-alpine AS build

# create a working directory inside the image
WORKDIR /app

# copy Go modules and dependencies to image
COPY . .

# RUN apk add --no-cache make

# download Go modules and dependencies
RUN go mod tidy

EXPOSE 3100

# compile application
RUN go build -o /binary

##
## STEP 2 - DEPLOY
##
FROM scratch

WORKDIR /

COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=build /binary /binary

EXPOSE 3100

# USER nonroot:nonroot

ENTRYPOINT ["/binary"]