export async function getDirectoryInfo(path: string) {
    let response = await fetch(`http://localhost:5000/api/directory?path=${path}`)
    let result = await response.json()
    return {
        Files: Object.keys(result.Files),
        Directories: Object.keys(result.Subdirectories)
    }
}

export async function getFileInfo(path: string) {
    let response = await fetch(`http://localhost:5000/api/file?path=${path}`)
    let name = path.split('/').pop()

    let data = await response.text()
    
    // download process
    var link = document.createElement('a')
    let contentType = `data:${(response.headers.get('Content-Type') || 'text/plain')}`
    console.log(data)
    link.href = `${contentType},${encodeURIComponent(data)}`
    link.download = name || 'file'
    link.click()
}