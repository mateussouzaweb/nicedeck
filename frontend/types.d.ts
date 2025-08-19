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
	id: number
	platform: string
	program: string
	layer: string
	type: string
	name: string
	description: string
	startDirectory: string
	executable: string
	launchOptions: string
	shortcutPath: string
	relativePath: string
	imagesPath: string
	iconPath: string
	iconUrl: string
	logoPath: string
	logoUrl: string
	coverPath: string
	coverUrl: string
	bannerPath: string
	bannerUrl: string
	heroPath: string
	heroUrl: string
	tags: string[]
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
	imagesPath: string
	steamRuntime: string
	steamPath: string
	steamAccountId: string
	steamAccountName: string
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
	id: number
}

interface LaunchShortcutResult {
	status: string
	error: string
}

interface ModifyShortcutData {
	action: string
	id: number
	platform: string
	program: string
	layer: string
	type: string
	name: string
	description: string
	startDirectory: string
	executable: string
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