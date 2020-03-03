export enum AppMessageType {
	UserMessage,
	UserReaction,
	GroupInvitation,
	SetGroupName,
}

export type UserMessage = {
	type: AppMessageType.UserMessage
	body: string
	attachments: Array<{ uri: string }>
}

export type UserReaction = {
	type: AppMessageType.UserReaction
	emoji: string
}

export type GroupInvitation = {
	type: AppMessageType.GroupInvitation
	groupPk: string
}

export type SetGroupName = {
	type: AppMessageType.SetGroupName
	name: string
}

export type AppMessage = UserMessage | UserReaction | GroupInvitation | SetGroupName
