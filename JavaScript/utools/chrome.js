function getChromeBookmarks() {
    let chromeDataDir = ''
    const profiles = ['Default', 'Profile 3', 'Profile 2', 'Profile 1']
    if (process.platform === 'win32') {
        chromeDataDir = path.join(process.env['LOCALAPPDATA'], 'Google/Chrome/User Data')
    } else if (process.platform === 'darwin') {
        chromeDataDir = path.join(window.utools.getPath('appData'), 'Google/Chrome')
    } else if (process.platform === 'linux') {
        chromeDataDir = path.join(window.utools.getPath('appData'), 'google-chrome')
    }
    const profile = profiles.find(profile => fs.existsSync(path.join(chromeDataDir, profile, 'Bookmarks')))
    if (!profile) return []
    const bookmarkPath = path.join(chromeDataDir, profile, 'Bookmarks')
    const bookmarksData = []
    try {
        const data = JSON.parse(fs.readFileSync(bookmarkPath, 'utf-8'))
        const getUrlData = (item) => {
            if (!item || !Array.isArray(item.children)) return
            item.children.forEach(c => {
                if (c.type === 'url') {
                    bookmarksData.push({
                        lowTitle: c.name.toLowerCase(),
                        title: c.name,
                        description: c.url,
                        icon: 'web.png'
                    })
                } else if (c.type === 'folder') {
                    getUrlData(c)
                }
            })
        }
        getUrlData(data.roots.bookmark_bar)
        getUrlData(data.roots.other)
        getUrlData(data.roots.synced)
    } catch (e) { }
    return bookmarksData
}