package utils

import (
	"net"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func RealIpFromFiberCtx(ctx *fiber.Ctx) string {
	val := ctx.Get("Cf-Connecting-Ip")
	if val != "" {
		if !IsLocalIp(val) {
			return val
		}
	}

	val = ctx.Get("Fastly-Client-Ip")
	if val != "" {
		if !IsLocalIp(val) {
			return val
		}
	}

	val = ctx.Get("True-Client-Ip")
	if val != "" {
		if !IsLocalIp(val) {
			return val
		}
	}

	val = ctx.Get("X-Real-IP")
	if val != "" {
		if !IsLocalIp(val) {
			return val
		}
	}

	val = ctx.Get("X-Client-IP")
	if val != "" {
		if !IsLocalIp(val) {
			return val
		}
	}

	val = ctx.Get("X-Original-Forwarded-For")
	if val != "" {
		for _, v := range strings.Split(val, ",") {
			if !IsLocalIp(v) {
				return v
			}
		}
		if !IsLocalIp(val) {
			return val
		}
	}

	val = ctx.Get("X-Forwarded-For")
	if val != "" {
		for _, v := range strings.Split(val, ",") {
			if net.ParseIP(v) != nil {
				return v
			}
		}
		if net.ParseIP(val) != nil {
			return val
		}
	}

	val = ctx.Get("X-Forwarded")
	if val != "" {
		for _, v := range strings.Split(val, ",") {
			if net.ParseIP(v) != nil {
				return v
			}
		}
		if net.ParseIP(val) != nil {
			return val
		}
	}

	val = ctx.Get("Forwarded-For")
	if val != "" {
		for _, v := range strings.Split(val, ",") {
			if net.ParseIP(v) != nil {
				return v
			}
		}
		if net.ParseIP(val) != nil {
			return val
		}
	}

	val = ctx.Get("Forwarded")
	if val != "" {
		for _, v := range strings.Split(val, ",") {
			if net.ParseIP(v) != nil {
				return v
			}
		}
		if net.ParseIP(val) != nil {
			return val
		}
	}

	for _, v := range ctx.IPs() {
		if net.ParseIP(v) != nil {
			return v
		}
	}

	return ctx.Context().RemoteIP().String()
}
