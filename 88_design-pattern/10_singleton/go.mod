module singleton

go 1.17

require (
	hungry v0.0.1
	lazy v0.0.1
)

replace (
	hungry => ./hungry
	lazy => ./lazy
)
