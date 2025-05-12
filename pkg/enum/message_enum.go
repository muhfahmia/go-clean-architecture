package enum

import "strings"

type MessageType string

const (
	// Success Messages (2XX)
	MessageCreatedSuccess   MessageType = "{Message} created successfully"
	MessageUpdatedSuccess   MessageType = "{Message} updated successfully"
	MessageDeletedSuccess   MessageType = "{Message} deleted successfully"
	MessageRetrievedSuccess MessageType = "{Message} data retrieved"
	MessageActionSuccess    MessageType = "Action completed successfully"
	MessageNoChanges        MessageType = "No changes detected ({Message} is up-to-date)"

	// Client Errors (4XX)
	MessageNotFound      MessageType = "{Message} not found"
	MessageAlreadyExists MessageType = "{Message} already exists"
	MessageInvalidData   MessageType = "Invalid {Message} data provided"
	MessageMissingFields MessageType = "Required fields are missing"
	MessageAccessDenied  MessageType = "Access denied for {Message}"
	MessageUnauthorized  MessageType = "Authentication required"
	MessageRequestLimit  MessageType = "Too many requests (try again later)"

	// Business Logic Errors
	MessageIncomplete      MessageType = "{Message} is incomplete (fill required fields)"
	MessageExpired         MessageType = "{Message} expired (renew to continue)"
	MessagePaymentRequired MessageType = "Payment required for {Message}"
	MessageInactive        MessageType = "{Message} is inactive"
	MessageConflict        MessageType = "{Message} conflict detected (duplicate or invalid state)"

	// Server Errors (5XX)
	MessageProcessFailed MessageType = "Failed to process {Message}"
)

// Helper function to replace placeholders
func (m MessageType) Format(msg string) string {
	return strings.ReplaceAll(string(m), "{Message}", msg)
}
