# Web-Screenshot-Plugin

A Web Screenshot plugin for [OpenAgents](https://openagents.com/).

# Introduction

This Web Screenshot Plugin utilize [Restpack Screenshot API](https://restpack.io/screenshot) that can capture screenshots of live web pages and deliver the results in JSON Format.  

[API Documentation](https://restpack.io/screenshot/docs).
# Build 

Include the library with Go get:

```bash
go get github.com/extism/go-pdk
```
You can find the reference documentation for this library on [pkg.go.dev](https://pkg.go.dev/github.com/extism/go-pdk).

# Usage
Run the plugin using this command :
```
extism call plugin.wasm run --wasi --allow-host="restpack.io" --input="WEB_URL"
```
Note : "WEB_URL" - The URL of web page, **including** the protocol that you want to capture.

## Usage Example & Screenshot Result
![Screenshot 2024-05-15 062520](https://github.com/nkn18/Web-Screenshot-Plugin/assets/57341598/641bac82-e5a9-45ab-9cbc-7c217b048467)
![Screenshot 2024-05-15 062559](https://github.com/nkn18/Web-Screenshot-Plugin/assets/57341598/fcf2adc3-ccef-435a-9352-142a237af4b4)
![Screenshot 2024-05-15 062614](https://github.com/nkn18/Web-Screenshot-Plugin/assets/57341598/d7a1184c-09c8-47e8-b486-595a464ff5fe)


