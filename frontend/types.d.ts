// Objects
interface Library {
	timestamp: number
	imagesPath: string
}

interface Program {
	id: string
	name: string
	description: string
	category: string
	tags: string[]
	folders: string[]
	flags: string[]
	website: string
	iconUrl: string
	logoUrl: string
	coverUrl: string
	bannerUrl: string
	heroUrl: string
}

interface ConsoleEmulator {
	name: string
	program: string
	extensions: string
	launchOptions: string
}

interface ConsolePlatform {
	name: string
	console: string
	folder: string
	emulators: ConsoleEmulator[]
}

interface NativePlatform {
	name: string
	runtime: string
	extensions: string
	startDirectory: string
	executable: string
	launchOptions: string
}

interface Shortcut {
	id: string
	program: string
	name: string
	description: string
	startDirectory: string
	executable: string
	launchOptions: string
	relativePath: string
	iconPath: string
	logoPath: string
	coverPath: string
	bannerPath: string
	heroPath: string
	tags: string[]
	timestamp: number
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
	data: Library
}

interface SaveLibraryResult {
	status: string
	error: string
}

interface SyncLibraryResult {
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
	console: ConsolePlatform[]
	native: NativePlatform[]
}

interface ListShortcutsResult {
	status: string
	error: string
	data: Shortcut[]
}

interface LaunchShortcutData {
	id: string
}

interface LaunchShortcutResult {
	status: string
	error: string
}

interface CreateShortcutData {
	name: string
	path: string
}

interface CreateShortcutResult {
	status: string
	error: string
}

interface AddShortcutData {
	id: string
	program: string
	name: string
	description: string
	startDirectory: string
	executable: string
	launchOptions: string
	relativePath: string
	iconPath: string
	logoPath: string
	coverPath: string
	bannerPath: string
	heroPath: string
	tags: string[]
}

interface AddShortcutResult {
	status: string
	error: string
}

interface ModifyShortcutData {
	action: string
	id: string
	program: string
	name: string
	description: string
	startDirectory: string
	executable: string
	launchOptions: string
	relativePath: string
	iconPath: string
	logoPath: string
	coverPath: string
	bannerPath: string
	heroPath: string
	tags: string[]
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