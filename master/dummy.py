from locust import Locust, TaskSet, task

class MyTaskSet(TaskSet):
    @task
    def my_task(self):
        print("executing my_task")

class MyLocust(Locust):
    task_set = MyTaskSet