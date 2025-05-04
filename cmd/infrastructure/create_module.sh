read -p "Enter your module name (PascalCase): " module_name

# Convert cases
module_lower=$(echo "$module_name" | tr '[:upper:]' '[:lower:]')
module_snake=$(echo "$module_name" | sed -E 's/([A-Z])/_\1/g' | sed 's/^_//' | tr '[:upper:]' '[:lower:]')
module_camel=$(echo "$module_name" | perl -pe 's/^([A-Z])/\l$1/')

# Base directory
BASE_DIR="internal"
mkdir -p "$BASE_DIR/entity"
mkdir -p "$BASE_DIR/repository"
mkdir -p "$BASE_DIR/usecase"
mkdir -p "$BASE_DIR/delivery/http"

# Generate Entity
cat > "$BASE_DIR/entity/${module_snake}_entity.go" <<EOF
package entity

type ${module_name}Entity struct {
    ${module_name}ID uint64 \`gorm:"primaryKey;autoIncrement;index:${module_snake}_idx"\`
    // Add other fields here
    Timestamp
}
EOF
echo "Created ${module_name} entity"

# Generate Repository
cat > "$BASE_DIR/repository/${module_snake}_repository.go" <<EOF
package repository

import (
	"github.com/muhfahmia/internal/entity"
	"gorm.io/gorm"
)

type ${module_name}Repository interface {
	// Add custom repository methods here
	Repository[entity.${module_name}Entity]
}

type ${module_camel}Repository struct {
	db *gorm.DB
	Repository[entity.${module_name}Entity]
}

func New${module_name}Repository(db *gorm.DB) ${module_name}Repository {
	return ${module_camel}Repository{
		db:         db,
		Repository: NewBaseRepository[entity.${module_name}Entity](db),
	}
}
EOF
echo "Created ${module_name} repository interfaces & implementation"

# Generate Usecase
cat > "$BASE_DIR/usecase/${module_snake}_usecase.go" <<EOF
package usecase

import (
	"github.com/muhfahmia/internal/repository"
)

type ${module_name}Usecase interface {
    Create() error
}

type ${module_camel}Usecase struct {
	repo repository.${module_name}Repository
}

func New${module_name}Usecase(repo repository.${module_name}Repository) ${module_name}Usecase {
	return &${module_camel}Usecase{repo: repo}
}

func (u *${module_camel}Usecase) Create() error {
	return nil
}
EOF
echo "Created ${module_name} usecase interfaces & implementation"

# Generate Fiber v2 HTTP Controller
cat > "$BASE_DIR/delivery/http/${module_snake}_controller.go" <<EOF
package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/muhfahmia/internal/usecase"
)

type ${module_name}Controller interface {
	Create(c *fiber.Ctx) error
}

type ${module_camel}Controller struct {
    usecase usecase.${module_name}Usecase
}

func New${module_name}Controller(usecase usecase.${module_name}Usecase) ${module_name}Controller {
	return &${module_camel}Controller{
		usecase:         usecase,
	}
}

func (s *${module_camel}Controller) Create(c *fiber.Ctx) error {
    return nil
}
EOF
echo "Created ${module_name} controller interfaces & implementation"

echo "Struktur clean architecture untuk entitas $module_name telah berhasil dibuat!"