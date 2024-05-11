interface Program {
    id: string
    name: string
    description: string
	category: string
    tags: string[]
    requiredFolders: string[]
    flatpakAppId: string
    flatpakOverrides: string[]
	flatpakArguments: string[]
    iconUrl: string
    logoUrl: string
    coverUrl: string
    bannerUrl: string
    heroUrl: string
}

interface Platform {
    name: string
	console: string
	emulator: string
	extensions: string
	launchOptions: string
}

interface Shortcut {
    appId: number     
	appName: string   
	startDir: string   
	exe: string 
	launchOptions: string  
	shortcutPath: string   
	icon: string  
	isHidden: number     
	allowDesktopConfig: number     
	allowOverlay: number     
	openVr: number    
	devkit: number    
	devkitGameId: string   
	devkitOverrideAppId: number     
	flatpakAppId: string   
	lastPlayTime: number     
	tags: string[]
	iconUrl: string
	logo: string
	logoUrl: string
	cover: string
	coverUrl: string
	banner: string
	bannerUrl: string
	hero: string
	heroUrl: string
	platform: string
	relativePath: string
}

interface ScrapeResult {
	name: string
	scrapeId: number
	shortcutId: number
	bannerUrls: string[]
	coverUrls: string[]
	heroUrls: string[]
	iconUrls: string[]
	logoUrls: string[]
}