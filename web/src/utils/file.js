// 读取文件
export const loadFile = (name) => {
  const xhr = new XMLHttpRequest()
  const okStatus = document.location.protocol === 'file:' ? 0 : 200
  xhr.open('GET', name, false)
  xhr.overrideMimeType('text/html;charset=utf-8')// 默认为utf-8
  xhr.send(null)
  return xhr.status === okStatus ? xhr.responseText : null
}

// unicode转utf-8
export const unicodeToUtf8 = (data) => {
  data = data.replace(/\\/g, '%')
  return unescape(data)
}

// 获取图片转base64
export const getBase64 = (file) => {
  return new Promise(function(resolve, reject) {
    const reader = new FileReader()
    let imgResult = ''
    reader.readAsDataURL(file)
    reader.onload = function() {
      imgResult = reader.result
    }
    reader.onerror = function(error) {
      reject(error)
    }
    reader.onloadend = function() {
      resolve(imgResult)
    }
  })
}
