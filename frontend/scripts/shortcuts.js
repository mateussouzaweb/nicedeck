// Shortcuts
window.addEventListener('load', () => {

    on('#shortcuts #load', 'click', async () => {
        const result = await request('GET', '/api/shortcuts')
        const data = JSON.parse(result || '[]')
        const items = data.map((item) => {
            return `<div class="item">
                <img loading="lazy" src="${item.coverUrl}" alt="${item.appName}" width="600" height="900"/><br/>
                <small>${item.appId}</small><br/>
                <h4>${item.appName}</h4>
                ${item.platform}
                ${item.platform ? '<br/>' : ''}
            </div>`
        })

        const list = $('#shortcuts #list')
        list.innerHTML = items.join('')
    })

})