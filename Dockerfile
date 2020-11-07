#------------------------------------
FROM golang:1.15-alpine AS build

WORKDIR /src

COPY . . 

RUN go build -o /out/app .

#------------------------------------
FROM scratch AS bin
COPY --from=build /out/app /

CMD [ "./app" ]