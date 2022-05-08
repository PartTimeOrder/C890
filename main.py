import time

import pyautogui

pyautogui.hotkey('command', 'a')

im1 = pyautogui.screenshot()
im1.save('my_screenshot.png')
print(im1)
