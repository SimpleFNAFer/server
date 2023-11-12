package home

type Repository interface {
	SaveRequest(sourceIP string)
	CountRequests() (count int)
	CountLastSecondRequests() (count int)
}
