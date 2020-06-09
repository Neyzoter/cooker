# coding: utf8

# # locust 版本1.0以前
# from locust import Locust, TaskSet, task
# class MyTaskSet(TaskSet):
#     @task(20)
#     def hello(self):
#         pass
#
#
# class Dummy(Locust):
#     task_set = MyTaskSet

# locust 版本1.0以后
from locust import User, task

class Dummy(User):
    @task(20)
    def hello(self):
        pass
