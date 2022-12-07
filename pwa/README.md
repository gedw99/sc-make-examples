# PWA


SO a notification ( for example ) will work as a PWA ( in browsers, and on desktop and mobile via the Webview), but the exact same code will also work on Desktop and Mobile without a Webview.

Backend: https://github.com/yomorun
- golang
- quic protocol
- edge based

webview
- we add any web assets inside
- we can then also add the testing automation robot into the webview to apply gestures, keysrokes and take screenshits / vnc of the running web app !!!


https://github.com/inkeliz/gowebview

webview http calls
https://github.com/inkeliz/go_inkwasm

Example fo calls to api and gui.
https://github.com/khmyznikov/ios-pwa-wrap/blob/main/pwa-shell/Info.plist#L113

testing
https://github.com/go-vgo/robotgo
- clicks the screen, does keyboard for desktop.
- mobile
    - https://github.com/nanoscopic/ios_remote_provider
    - https://github.com/danielpaulus/go-ios

https://github.com/go-rod/rod
- automation and scrape usin web.


sys tray
https://github.com/getlantern/systray
