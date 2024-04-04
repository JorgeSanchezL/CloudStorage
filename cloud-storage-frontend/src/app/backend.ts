export async function getDirectoryInfo(path: string) {
    let response = await fetch(`http://localhost:5000/api/directory?path=${path}`)
    let result = await response.json()
    return {
        Files: Object.keys(result.Files),
        Directories: Object.keys(result.Subdirectories)
    }
}