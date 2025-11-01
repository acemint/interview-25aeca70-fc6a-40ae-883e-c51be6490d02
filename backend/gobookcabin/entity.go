package gobookcabin

type BaseEntity struct {
	ID        int     `gorm:"primaryKey,autoIncrement,<-:false"`
	CreatedAt *string `gorm:"<-:false"`
}
