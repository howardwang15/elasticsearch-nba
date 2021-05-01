from selenium import webdriver
from selenium.webdriver.chrome.options import Options
from pyvirtualdisplay import Display
import os
import pprint
import time
import json

display = Display(visible=0, size=(800, 600))
display.start()

os.environ["LANG"] = "en_US.UTF-8"
DRIVER_PATH = 'chromedriver'

options = Options()
options.headless = True
options.add_argument("--window-size=1920,1200")
options.add_argument("--disable-dev-shm-usage")
options.add_argument("--no-sandbox")

driver = webdriver.Chrome(options=options, executable_path=DRIVER_PATH)
driver.get('https://www.espn.com/nba/stats/player/_/table')

NUM_ATTRIBUTES = 20
attributes = [
    'GP',
    'MIN',
    'PTS',
    'FGM',
    'FGA',
    'FG%',
    '3PM',
    '3PA',
    '3P%',
    'FTM',
    'FTA',
    'FT%',
    'REB',
    'AST',
    'STL',
    'BLK',
    'TO',
    'DD2',
    'TD3',
    'PER'
]

while True:
    try:
        driver.find_element_by_css_selector('a.AnchorLink.loadMore__link').click()
        time.sleep(2)
    except Exception:
        break

players_table = driver.find_element_by_css_selector('.Table--fixed-left.Table--align-right')
players = list(map(lambda p: p.get_attribute('innerHTML'),
    players_table.find_elements_by_css_selector('a.AnchorLink')))

stats_table = driver.find_elements_by_css_selector('tbody.Table__TBODY')[1]
stats = list(map(lambda p: p.get_attribute('innerHTML'),
    stats_table.find_elements_by_css_selector('td.Table__TD')))

stats = [float(stat) for i, stat in enumerate(stats) if i % (NUM_ATTRIBUTES+1) != 0]
stats = [stats[i:(i+NUM_ATTRIBUTES+1)] for i in range(0, len(stats), NUM_ATTRIBUTES)]

data = []
for player, stat in zip(players, stats):
    info = {}
    info['name'] = player
    for attribute, s in zip(attributes, stat):
        info[attribute] = s
    data.append(json.dumps(info))

with open('/data/players.json', 'w') as f:
    f.write('\n'.join(data))
print(data)
driver.close()
