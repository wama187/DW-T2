package book_v2
import (
	"context"
	"time"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"app/server/apps/book"
)

type bookRepoMongo struct {
	coll *mongo.Collection
	ctx  context.Context
}

func NewBookRepositoryMongo(client *mongo.Client, dbName, collName string) *bookRepoMongo {
	return &bookRepoMongo{
		coll: client.Database(dbName).Collection(collName),
		ctx:  context.Background(),
	}
}

func (r *bookRepoMongo) Create(b *book.BookMongo) error {
	doc := bson.M{
		"title":      b.Title,
		"author":     b.Author,
		"owner_id":   b.OwnerID,
		"created_at": time.Now(),
		"updated_at": time.Now(),
	}
	_, err := r.coll.InsertOne(r.ctx, doc)
	return err
}

func (r *bookRepoMongo) Update(b *book.BookMongo) error {
    update := bson.M{
        "$set": bson.M{
            "title":      b.Title,
            "author":     b.Author,
            "updated_at": time.Now(),
        },
    }

    id, err := primitive.ObjectIDFromHex(b.ID)
    if err != nil {
        return fmt.Errorf("ID inválido: %w", err)
    }

    _, err = r.coll.UpdateOne(r.ctx, bson.M{"_id": id, "owner_id": b.OwnerID}, update)
    return err
}


func (r *bookRepoMongo) Delete(bookID string, ownerID uint) error {
	id, err := primitive.ObjectIDFromHex(bookID)
    if err != nil {
        return fmt.Errorf("ID inválido: %w", err)
    }
	_, err = r.coll.DeleteOne(r.ctx, bson.M{"_id": id, "owner_id": ownerID})
	return err
}

func (r *bookRepoMongo) FindByID(b *book.BookMongo) (*book.BookMongo, error) {
	var doc bson.M

	id, err := primitive.ObjectIDFromHex(b.ID)
    if err != nil {
        return nil, fmt.Errorf("ID inválido: %w", err)
    }

	err = r.coll.FindOne(r.ctx, bson.M{"_id": id, "owner_id": b.OwnerID}).Decode(&doc)
	if err != nil {
		return nil, err
	}

	return &book.BookMongo{
		ID:      doc["_id"].(primitive.ObjectID).Hex(),
		Title:   doc["title"].(string),
		Author:  doc["author"].(string),
		OwnerID: uint(doc["owner_id"].(int64)),
	}, nil
}


func (r *bookRepoMongo) FindAllByOwner(ownerID uint) ([]*book.BookMongo, error) {
    cursor, err := r.coll.Find(r.ctx, bson.M{"owner_id": ownerID})
    if err != nil {
        return nil, err
    }
    defer cursor.Close(r.ctx)

    var books []*book.BookMongo
    for cursor.Next(r.ctx) {
        var doc bson.M
        if err := cursor.Decode(&doc); err != nil {
            return nil, err
        }

        books = append(books, &book.BookMongo{
            ID:      doc["_id"].(primitive.ObjectID).Hex(),
            Title:   doc["title"].(string),
            Author:  doc["author"].(string),
            OwnerID: uint(doc["owner_id"].(int64)),
        })
    }
    return books, nil
}

