// Platforms
window.addEventListener('load', async () => {

    /**
     * Load and show available platforms in the software
     */
    async function loadPlatforms() {

        /** @type {ListPlatformsResult} */
        const request = await requestJson('GET', '/api/platforms')
        const html = []

        const toKey = (name) => {
            return name.toLowerCase().replaceAll(' ', '-')
        }

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

        append('Consoles', request.console.filter((platform) => {
            return platform.available
        }).map((platform) => {
            const key = toKey(platform.console)
            const isCustom = platform.type === 'custom'
            return {
                key: key,
                name: platform.console,
                value: platform.name,
                image: `/img/platforms/${isCustom ? 'custom' : key}.png`
            }
        }))

        append('Launchers', request.launcher.filter((platform) => {
            return platform.available
        }).map((platform) => {
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