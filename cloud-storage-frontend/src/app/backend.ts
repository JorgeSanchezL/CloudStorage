type FileInfo = {
    Name: string;
    Size: number;
    LastModification: string;
};

export async function getDirectoryInfo(path: string) {
    let response = await fetch(`http://localhost:5000/api/directory?path=${path}`)
    let result = await response.json()
    return {
        Files: Object.keys(result.Files).map(key => ({
            Name: key,
            Size: result.Files[key].Size,
            LastModification: result.Files[key].LastModification
        })),
        Directories: Object.keys(result.Subdirectories) // We could be doing the same thing here, but we only need to store the names
    }
}

export async function getFile(path: string) {
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