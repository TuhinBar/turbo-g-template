package services

import (
	"context"
	"time"
	database "turbo-g-template/server/configs"
	"turbo-g-template/server/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetAllTasks() ([]models.Task, error) {
    collection := database.Client.Database("taskdb").Collection("tasks")
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    var tasks []models.Task
    cursor, err := collection.Find(ctx, bson.M{})
    if err != nil {
        return nil, err
    }
    defer cursor.Close(ctx)

    if err = cursor.All(ctx, &tasks); err != nil {
        return nil, err
    }

    return tasks, nil
}

func CreateTask(task *models.Task) (*models.Task, error) {
    collection := database.Client.Database("taskdb").Collection("tasks")
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    task.CreatedAt = time.Now()
    task.UpdatedAt = time.Now()

    result, err := collection.InsertOne(ctx, task)
    if err != nil {
        return nil, err
    }

    task.ID = result.InsertedID.(primitive.ObjectID)
    return task, nil
}
