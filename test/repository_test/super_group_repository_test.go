package repository_test

// import (
// 	"testing"

// 	"a2sv.org/hub/Domain/entity"

// 	"github.com/stretchr/testify/assert"
// 	"gorm.io/gorm"
// 	"github.com/DATA-DOG/go-sqlmock"
// 	"gorm.io/driver/postgres"
// 	repository"a2sv.org/hub/Repository/postgres"
	
// )

// func setupTestDB(t *testing.T) *gorm.DB {
// 	// Use postgres mock instead of sqlite since CGO is disabled
// 	mockDB, mock, err := sqlmock.New()
// 	if err != nil {
// 		t.Fatalf("Failed to create mock DB: %v", err)
// 	}

// 	dialector := postgres.New(postgres.Config{
// 		Conn: mockDB,
// 	})

// 	db, err := gorm.Open(dialector, &gorm.Config{})
// 	if err != nil {
// 		t.Fatalf("Failed to open mock database: %v", err)
// 	}

// 	// Set up expected migrations
// 	mock.ExpectExec("CREATE TABLE").WillReturnResult(sqlmock.NewResult(0, 0))

// 	err = db.AutoMigrate(&entity.SuperGroup{}, &entity.Group{}, &entity.SuperToGroup{})
// 	if err != nil {
// 		t.Fatalf("Failed to migrate database: %v", err)
// 	}

// 	return db
// }

// func TestSuperGroupRepository_Create(t *testing.T) {
// 	db := setupTestDB(t)
// 	repo := repository.NewSuperGroupRepository(db)

// 	t.Run("successful creation", func(t *testing.T) {
// 		// Set up mock expectations
// 		mock, err := db.DB()
// 		if err != nil {
// 			t.Fatalf("Failed to get mock DB: %v", err)
// 		}
// 		sqlMock := mock.(*sql.DB)
// 		_, ok := sqlMock.(sqlmock.Sqlmock) 
// 		if !ok {
// 			t.Fatal("Failed to get SQL mock")
// 		}

// 		sqlMock.ExpectBegin()
// 		sqlMock.ExpectQuery(`INSERT INTO "super_groups"`).
// 			WillReturnRows(sqlmock.NewRows([]string{"id", "name"}).
// 				AddRow(1, "Test Super Group"))
// 		sqlMock.ExpectCommit()

// 		name := "Test Super Group"
// 		superGroup, err := repo.Create(name)

// 		assert.NoError(t, err)
// 		assert.NotNil(t, superGroup)
// 		assert.Equal(t, name, superGroup.Name)
// 		assert.NotZero(t, superGroup.ID)

// 		// Verify all expectations were met
// 		if err := sqlMock.ExpectationsWereMet(); err != nil {
// 			t.Errorf("there were unfulfilled expectations: %s", err)
// 		}
// 	})
// }

// func TestSuperGroupRepository_GetByID(t *testing.T) {
// 	db := setupTestDB(t)
// 	repo := repository.NewSuperGroupRepository(db)

// 	t.Run("get existing super group", func(t *testing.T) {
// 		// Create test data
// 		name := "Test Super Group"
// 		created, _ := repo.Create(name)

// 		// Test retrieval
// 		found, err := repo.GetByID(created.ID)
// 		assert.NoError(t, err)
// 		assert.Equal(t, created.ID, found.ID)
// 		assert.Equal(t, name, found.Name)
// 	})

// 	t.Run("get non-existent super group", func(t *testing.T) {
// 		found, err := repo.GetByID(9999)
// 		assert.Error(t, err)
// 		assert.Nil(t, found)
// 	})
// }

// func TestSuperGroupRepository_Update(t *testing.T) {
// 	db := setupTestDB(t)
// 	repo := repository.NewSuperGroupRepository(db)

// 	t.Run("update existing super group", func(t *testing.T) {
// 		// Create test data
// 		original, _ := repo.Create("Original Name")

// 		// Update
// 		newName := "Updated Name"
// 		updated, err := repo.Update(original.ID, newName)

// 		assert.NoError(t, err)
// 		assert.Equal(t, original.ID, updated.ID)
// 		assert.Equal(t, newName, updated.Name)
// 	})
// }

// func TestSuperGroupRepository_Delete(t *testing.T) {
// 	db := setupTestDB(t)
// 	repo := repository.NewSuperGroupRepository(db)

// 	t.Run("delete existing super group", func(t *testing.T) {
// 		// Create test data
// 		superGroup, _ := repo.Create("Test Super Group")

// 		// Delete
// 		err := repo.Delete(superGroup.ID)
// 		assert.NoError(t, err)

// 		// Verify deletion
// 		found, err := repo.GetByID(superGroup.ID)
// 		assert.Error(t, err)
// 		assert.Nil(t, found)
// 	})
// }

// func TestSuperGroupRepository_List(t *testing.T) {
// 	db := setupTestDB(t)
// 	repo := repository.NewSuperGroupRepository(db)

// 	t.Run("list super groups", func(t *testing.T) {
// 		// Create test data
// 		repo.Create("Super Group 1")
// 		repo.Create("Super Group 2")

// 		// List all
// 		superGroups, err := repo.List()
// 		assert.NoError(t, err)
// 		assert.Len(t, superGroups, 2)
// 	})
// }

// func TestSuperGroupRepository_GroupOperations(t *testing.T) {
// 	db := setupTestDB(t)
// 	repo := repository.NewSuperGroupRepository(db)

// 	t.Run("add and remove groups", func(t *testing.T) {
// 		// Create test data
// 		superGroup, _ := repo.Create("Test Super Group")
		
// 		// Create test groups
// 		group1 := &entity.Group{Name: "Group 1"}
// 		group2 := &entity.Group{Name: "Group 2"}
// 		db.Create(group1)
// 		db.Create(group2)

// 		// Add groups
// 		err := repo.AddGroup(superGroup.ID, group1.ID)
// 		assert.NoError(t, err)
// 		err = repo.AddGroup(superGroup.ID, group2.ID)
// 		assert.NoError(t, err)

// 		// Get groups
// 		groups, err := repo.GetGroups(superGroup.ID)
// 		assert.NoError(t, err)
// 		assert.Len(t, groups, 2)

// 		// Remove group
// 		err = repo.RemoveGroup(superGroup.ID, group1.ID)
// 		assert.NoError(t, err)

// 		// Verify removal
// 		groups, err = repo.GetGroups(superGroup.ID)
// 		assert.NoError(t, err)
// 		assert.Len(t, groups, 1)
// 		assert.Equal(t, group2.ID, groups[0].ID)
// 	})
// }
