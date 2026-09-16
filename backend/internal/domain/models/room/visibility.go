package room

type Visibility string

const (
	VisibilityPublic      Visibility = "public"
	VisibilityFriendsOnly Visibility = "frinends_only"
	VisibilityPrivate     Visibility = "Private"
)

func (v Visibility) isValid() bool {
	switch v {
	case VisibilityPublic, VisibilityFriendsOnly, VisibilityPrivate:
		return true
	default:
		return false
	}
}
