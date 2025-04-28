package repository
import(
	"a2sv.org/hub/Domain/entity"
)


// ProblemTracksRepository defines methods for Problem Tracks data operations

type ProblemInTracksRepository interface {
	AddProblemToTrack(trackID uint,problemID uint) error

	ListProblemsInTrack(trackID uint) ([]*entity.Problem, error)

	GetProblemInTracksByName(trackID uint,name string) (*entity.Problem, error)
	GetProblemInTracksByDifficulty(trackID uint,difficulty string) ([]*entity.Problem, error)
	GetProblemInTracksByTag(trackID uint,tag string) ([]*entity.Problem, error)
	GetProblemInTracksByPlatform(trackID uint,platform string) ([]*entity.Problem, error)

	RemoveProblemFromTrack(id uint) error
}
