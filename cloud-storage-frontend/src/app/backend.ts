export async function getDirectoryInfo(path: string) {
    let response = await fetch(`http://localhost:5000/api/directory?path=${path}`)
    let result = await response.json()
    return {
        Files: Object.keys(result.Files),
        Directories: Object.keys(result.Subdirectories)
    }
}

export async function getFileInfo(path: string) {
    fetch(`http://localhost:5000/api/file?path=${path}`)
            .then(response => {
                response.blob().then(blob => {
                    let url = window.URL.createObjectURL(blob)
                    let a = document.createElement('a')
                    let name = path.split('/').pop()
                    a.href = url;
                    a.download = name || 'file'
                    a.click()
                });
        });
}