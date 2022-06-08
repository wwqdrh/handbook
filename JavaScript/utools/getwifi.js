const cp = require('child_process')
let wifiPassword
if (utools.isWindows()) {
  let stdoutBuffer = cp.execFileSync('netsh', ['wlan', 'show', 'interface'])
  let stdoutText = new TextDecoder('gbk').decode(stdoutBuffer)
  let ret = /^\s*SSID\s*: (.+)\s*$/gm.exec(stdoutText)
  if (!ret) {
    return utools.showNotification('未找到网络已连接的 Wi-Fi 名称')
  }
  const args = ['wlan', 'show', 'profile', `name=${ret[1]}`, 'key=clear']
  stdoutBuffer = cp.execFileSync('netsh', args)
  stdoutText = new TextDecoder('gbk').decode(stdoutBuffer)
  ret = /^\s*(?:Key Content|关键内容)\s*: (.+)\s*$/gm.exec(stdoutText)
  if (!ret) {
    return utools.showNotification('未能获取 Wi-Fi 密码')
  }
  wifiPassword = ret[1]
} else if (utools.isMacOs()) {
  let stdoutText = cp.execFileSync('/System/Library/PrivateFrameworks/Apple80211.framework/Versions/Current/Resources/airport', ['-I']).toString()
	if (stdoutText.includes('AirPort: Off')) {
		return utools.showNotification('Wi-Fi 已关闭')
	}
	let ret = /^\s*SSID: (.+)\s*$/gm.exec(stdoutText)
  if (!ret) {
    return utools.showNotification('未找到网络已连接的 Wi-Fi 名称')
  }
	const args = ['find-generic-password', '-D', 'AirPort network password', '-wa', ret[1]]
  try {
    wifiPassword = cp.execFileSync('security', args).toString().trim()
  } catch (__) {}
  if (!wifiPassword) return
}
utools.copyText(wifiPassword)
utools.showNotification(`Wi-Fi 密码 "${wifiPassword}" 已复制`)
