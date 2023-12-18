package solid

type EmailService struct {
}

func (e *EmailService) Send(to string, message string) {
	// send email
}

type NotificationService struct {
	emailService *EmailService
}

func (n *NotificationService) Notify(to string, message string) {
	n.emailService.Send(to, message)
}

// In this case, the `NotificationService` directly depends on the `EmailService`, making it difficult to switch to another notification method (e.g., SMS) or test the `NotificationService` in isolation. To follow the Dependency Inversion Principle, we can introduce an interface and depend on that instead:

type MessageSender interface {
	Send(to string, message string)
}

type EmailService1 struct {
}

func (e *EmailService1) Send(to string, message string) {
	// send email
}

type SMSService struct {
}

func (s *SMSService) Send(to string, message string) {
	// send SMS
}

type NotificationService1 struct {
	MessageSender MessageSender
}

func (n *NotificationService1) Notify(to string, message string) {
	n.MessageSender.Send(to, message)
}
