// Objects
interface Program {
	id: string
	name: string
	description: string
	category: string
	tags: string[]
	folders: string[]
	website: string
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

// Requests
interface LoadLibraryResult {
	status: string
	error: string
	steamRuntime: string
	steamPath: string
	configPath: string
	artworksPath: string
}

interface SaveLibraryResult {
	status: string
	error: string
}

interface ListProgramsResult {
	status: string
	error: string
	data: Program[]
}

interface ListPlatformsResult {
	status: string
	error: string
	data: Platform[]
}

interface ListShortcutsResult {
	status: string
	error: string
	data: Shortcut[]
}

interface LaunchShortcutData {
	appId: number
}

interface LaunchShortcutResult {
	status: string
	error: string
}

interface ModifyShortcutData {
	action: string
	appId: number
	appName: string
	startDir: string
	exe: string
	launchOptions: string
	iconUrl: string
	logoUrl: string
	coverUrl: string
	bannerUrl: string
	heroUrl: string
}

interface ModifyShortcutResult {
	status: string
	error: string
}

interface InstallProgramsData {
	programs: string[]
}

interface InstallProgramsResult {
	status: string
	error: string
}

interface RemoveProgramsData {
	programs: string[]
}

interface RemoveProgramsResult {
	status: string
	error: string
}

interface BackupStateData {
	platforms: string[]
	preferences: string[]
}

interface BackupStateResult {
	status: string
	error: string
}

interface RestoreStateData {
	platforms: string[]
	preferences: string[]
}

interface RestoreStateResult {
	status: string
	error: string
}

interface ProcessROMsData {
	platforms: string[]
	preferences: string[]
}

interface ProcessROMsResult {
	status: string
	error: string
}

interface ScrapeDataResult {
	status: string
	error: string
	result: ScrapeResult
}

interface OpenLinkData {
	link: string
}

interface OpenLinkResult {
	status: string
	error: string
}