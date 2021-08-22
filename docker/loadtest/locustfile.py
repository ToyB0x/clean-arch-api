import time

from locust import HttpUser, task, between
import json
import random
import uuid

# NOTE: キャッシュシステム利用のシナリオ
class ApiUser(HttpUser):
    wait_time = between(1.0, 2.0)

   # 2/3のユーザは2ヶ月分のカレンダーを表示するだけで予約しない
    @task(2)
    def check_calendar(self):
        self.client.get('/schedules-memstore/date/2000/1')
        # 3秒後に翌月を表示
        time.sleep(3)
        self.client.get('/schedules-memstore/date/2000/2')

    # 1/3のユーザは2ヶ月分のカレンダーを表示した上で空枠が有れば予約
    @task(1)
    def check_and_reserve(self):
        calendar1 = self.client.get('/schedules-memstore/date/2000/1')
        # 3秒後に翌月を表示
        time.sleep(3)
        calendar2 = self.client.get('/schedules-memstore/date/2000/2')
        # 5秒後に予約登録
        time.sleep(5)
        available_dates = find_available_date([calendar1, calendar2])
        if available_dates:
            selected_schedule = random.choice(available_dates)
            ticket_id = str(uuid.uuid4())
            payload = {
                "ticket_id": ticket_id,
                "schedule_id": selected_schedule["id"],
            }

            headers = {'content-type': 'application/json'}
            url = '/reservations'
            self.client.post(url, data=json.dumps(payload), headers=headers)

# NOTE: キャッシュシステム未利用のシナリオ
# class ApiUser(HttpUser):
#     wait_time = between(1.0, 2.0)
#
#     @task(2)
#     def check_calendar(self):
#         self.client.get('/schedules/date/2000/1')
#         time.sleep(3)
#         self.client.get('/schedules/date/2000/2')
#
#     @task(1)
#     def check_and_reserve(self):
#         calendar1 = self.client.get('/schedules/date/2000/1')
#         time.sleep(3)
#         calendar2 = self.client.get('/schedules/date/2000/2')
#         time.sleep(5)
#         available_dates = find_available_date([calendar1, calendar2])
#         if available_dates:
#             selected_schedule = random.choice(available_dates)
#             ticket_id = str(uuid.uuid4())
#             payload = {
#                 "ticket_id": ticket_id,
#                 "schedule_id": selected_schedule["id"],
#             }
#
#             headers = {'content-type': 'application/json'}
#             url = '/reservations'
#             self.client.post(url, data=json.dumps(payload), headers=headers)

# utils
def find_available_date(calendars):
    available_dates = []
    for calendar in calendars:
        schedules = calendar.json()
        for schedule in schedules:
            if schedule["stock"] > 0:
                available_dates.append(schedule)
    return available_dates
