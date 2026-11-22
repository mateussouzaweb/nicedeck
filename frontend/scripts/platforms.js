// Platforms
window.addEventListener('load', async () => {

    /** @type {ConsolePlatform[]} */
    let consoles = []

    /** @type {LaunchersPlatform[]} */
    let launchers = []

    /** @type {NativePlatform[]} */
    let natives = []

    /**
     * Retrieve console platform by name
     * @param {String} name
     * @returns {ConsolePlatform}
     */
    function getConsole(name) {
        return consoles.find((item) => String(item.name) === String(name))
    }

    /**
     * Retrieve launcher platform by name
     * @param {String} name
     * @returns {LauncherPlatform}
     */
    function getLauncher(name) {
        return launchers.find((item) => String(item.name) === String(name))
    }

    /**
     * Retrieve native platform by name
     * @param {String} name
     * @returns {NativePlatform}
     */
    function getNative(name) {
        return natives.find((item) => String(item.name) === String(name))
    }

    /**
     * Turn name into consistent key
     * @param {String} name
     * @returns {String}
     */
    const toKey = (name) => {
        return name.toLowerCase().replaceAll(' ', '-')
    }

    /**
     * Render platforms in HTML
     */
    async function renderPlatforms() {

        const html = []
        const append = (category, platforms) => {
            if (platforms.length === 0) {
                return
            }

            html.push(
            `<section class="group group-${category.toLowerCase()}">
                <div class="group-info">
                    <h4>${category}</h4>
                    <p class="mass-actions">
                        <span class="select-all">SELECT ALL</span>
                        <span class="separator">/</span>
                        <span class="clear-all">CLEAR</span>
                    </p>
                </div>`)

            platforms.map((platform) => {
                html.push( `<label class="checkbox" title="${platform.name}">
                    <input type="checkbox" name="platforms[]" value="${platform.value}" checked="checked" />
                    <div class="area">
                        <div class="icon">
                            <img loading="lazy" src="${platform.image}" alt="${platform.name}" width="96" height="96" />
                        </div>
                        <div class="info">
                            <b>${platform.value}</b>
                        </div>
                        <div class="details" data-details-platform="${platform.value}" title="Details">
                            <img src="./img/icons/info.svg" alt="Details" width="20" height="20" />
                        </div>
                    </div>
                </label>`
                )
            })

            html.push(`<div class="fill"></div>`)
            html.push(`<div class="fill"></div>`)
            html.push(`<div class="fill"></div>`)
            html.push(`<div class="fill"></div>`)
            html.push(`<div class="fill"></div>`)
            html.push(`<div class="fill"></div>`)
            html.push(`<div class="fill"></div>`)
            html.push(`<div class="fill"></div>`)
            html.push(`</section>`)
        }

        append('Consoles', consoles.map((platform) => {
            const key = toKey(platform.console)
            const isCustom = platform.type === 'custom'
            return {
                key: key,
                name: platform.console,
                value: platform.name,
                image: `/img/platforms/${isCustom ? 'custom' : key}.png`
            }
        }))

        append('Launchers', launchers.map((platform) => {
            return {
                key: toKey(platform.launcher),
                name: platform.launcher,
                value: platform.name,
                image: `/img/programs/${platform.program}.png`
            }
        }))

        const destination = $('#platforms .platforms')
        destination.innerHTML = html.join('')

    }

    /**
     * Render platform details modal
     * @param {String} name
     */
    async function renderPlatformDetails(name) {

        const modal = $('#modal-details-platform')
        const content = $('.content', modal)
        const html = []

        const console = getConsole(name)
        const launcher = getLauncher(name)
        const native = getNative(name)

        if( native ){

            html.push(`
                <div class="columns">
                    <div class="info">
                        <h4><b>${native.name}</b></h4>
                        <span>Installed:</span> ${native.installed ? 'Yes': 'No'}<br/>
                        <span>Runtime:</span> ${native.runtime}<br/>
                        <span>Extensions:</span> <span class="tag">${native.extensions}</span>
                    </div>
                </div>
            `)

        } else if( launcher ){

            const image = `/img/programs/${launcher.program}.png`

            html.push(`
                <div class="columns">
                    <div class="icon">
                        <img loading="lazy" src="${image}" alt="${launcher.name}" width="96" height="96" />
                    </div>
                    <div class="info">
                        <h4><b>${launcher.name}</b></h4>
                        <span>Launcher:</span> ${launcher.launcher}<br/>
                        <span>Installed:</span> ${launcher.installed ? 'Yes': 'No'}
                    </div>
                </div>
            `)

        } else if( console ){

            const key = toKey(console.console)
            const isCustom = console.type === 'custom'
            const image = `/img/platforms/${isCustom ? 'custom' : key}.png`
            const folder = `$ROMS/${console.folder}`

            const emulators = console.emulators.map((emulator) => {
                return `<div class="emulator">
                    <span>Emulator:</span> ${emulator.name}<br/>
                    <span>Installed:</span> ${emulator.installed ? 'Yes': 'No'}<br/>
                    <span>Folders:</span> <span class="tag">${folder}</span> <span class="tag">${folder}/${emulator.name}</span><br/>
                    <span>Extensions:</span> <span class="tag">${emulator.extensions}</span><br/>
                    <span>Launch Format:</span> <span class="tag">${emulator.launchOptions}</span>
                </div>`
            }).join('')

            html.push(`
                <div class="columns">
                    <div class="icon">
                        <img loading="lazy" src="${image}" alt="${console.name}" width="96" height="96" />
                    </div>
                    <div class="info">
                        <h4><b>${console.name}</b></h4>
                        <span>Console:</span> ${console.console}<br/>
                        <span>Type:</span> ${isCustom ? 'Custom' : 'Native'}<br/>
                        <span>Folder:</span> <span class="tag">$ROMS</span>
                    </div>
                </div>
                <div class="emulators">
                    ${emulators}
                </div>
            `)

        }

        modal.dataset.platform = name
        content.innerHTML = html.join('')
        window.showModal(modal)

    }

    /**
     * Load and show available platforms in the software
     */
    async function loadPlatforms() {

        /** @type {ListPlatformsResult} */
        const request = await requestJson('GET', '/api/platforms')

        consoles = request.console.filter((platform) => {
            return platform.available
        })

        launchers = request.launcher.filter((platform) => {
            return platform.available
        })

        natives = request.native.filter((platform) => {
            return platform.available
        })

        await renderPlatforms()

    }

    on('#platforms [data-details-platform]', 'click', async (event) => {
        event.preventDefault()

        const button = event.target.closest('[data-details-platform]')
        if (button.disabled) {
            return
        }

        const name = button.dataset.detailsPlatform
        await renderPlatformDetails(name)
    })

    on('#platforms .select-all', 'click', (event) => {
        event.preventDefault()
        const parent = event.target.closest('.group')
        const inputs = $$('input[type="checkbox"]', parent)
        inputs.map((input) => input.checked = "checked")
    })

    on('#platforms .clear-all', 'click', (event) => {
        event.preventDefault()
        const parent = event.target.closest('.group')
        const inputs = $$('input[type="checkbox"]', parent)
        inputs.map((input) => input.checked = "")
    })

    on('#platforms form', 'submit', async (event) => {
        event.preventDefault()

        const form = $('#platforms form')
        const button = $('button[type="submit"]', form)

        if (button.disabled) {
            return
        }

        const data = new FormData(form)

        /** @type {ProcessPlatformsData} */
        const body = {
            platforms: data.getAll('platforms[]'),
            preferences: data.getAll('preferences[]')
        }

        await window.runAndCaptureConsole(button, true, async () => {
            try {
                /** @type {LoadLibraryResult} */
                await requestJson('POST', '/api/library/load')
                /** @type {ProcessPlatformsResult} */
                await requestJson('POST', '/api/platforms/process', JSON.stringify(body))
            } catch (error) {
                window.showError(error)
            }

            try {
                /** @type {SaveLibraryResult} */
                await requestJson('POST', '/api/library/save')
            } catch (error) {
                window.showError(error)
            }
        })

        try {
            await window.reloadShortcuts()
        } catch (error) {
            window.showError(error)
        }
    })

    try {
        await loadPlatforms()
    } catch (error) {
        window.showError(error)
    }

    window.reloadPlatforms = loadPlatforms

})