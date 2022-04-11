package utils

import (
	"github.com/gofiber/fiber/v2/middleware/logger"
	"os"
)

var file, _ = os.OpenFile("./logs/kma_score.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0755)

var LogToFile = logger.New(logger.Config{
	Format:     "[${time}]: ${ip} - ${method} ${status} ${path} ${latency} ${bytesSent}B\n",
	TimeFormat: "2006-01-02 15:04:05",
	TimeZone:   "Asia/Bangkok",
	Output:     file,
})

var LogToTerminal = logger.New(logger.Config{
	Format:     "[${time}]: ${ip} - ${method} ${status} ${path} ${latency} ${bytesSent}B\n",
	TimeFormat: "2006-01-02 15:04:05",
	TimeZone:   "Asia/Bangkok",
})
