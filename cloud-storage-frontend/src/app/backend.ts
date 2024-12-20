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

export async function removeFile(path: string) {
    fetch(`http://localhost:5000/api/file?path=${path}`, { method: 'DELETE' })
        .then(response => {
            if (response.ok) {
                console.log('File deleted successfully');
            } else {
                console.error('File deletion failed');
            }
        });
}

export async function uploadFiles(files: FileList, paths: string[]) {
    const formData = new FormData();
    paths.forEach(path => {
        formData.append('paths', path);
    })
    for (let i = 0; i < files.length; i++) {
        formData.append('files', files.item(i) as Blob);
    }

    try {
        const response = await fetch('http://localhost:5000/api/file', {
            method: 'POST',
            body: formData
        });

        if (response.ok) {
            console.log('File uploaded successfully');
        } else {
            console.error('File upload failed');
        }
    } catch (error) {
        console.error('Error uploading file:', error);
    }
}