#!/usr/bin/env python3
import collections
import os.path
import sys

DGRAM_MAX_D_VALUE = 3

class PatternSearch2D():
    def __init__(self, pattern: list[str]):
        if not pattern:
            raise ValueError('pattern can not be empty string')

        self.__pattern_matrix = pattern
        self.__pattern_min_row_length = min([len(row) for row in self.__pattern_matrix])

        self.__d_value = max(min(DGRAM_MAX_D_VALUE, self.__pattern_min_row_length - 1), 1)
        self.__dgram_vertical_skip = collections.defaultdict(lambda: len(self.__pattern_matrix))

        self.__dgram_hor_pos = collections.defaultdict(lambda: -1)
        self.__dgram_hor_pos_linked_list = [0] * self.__pattern_min_row_length
        self.__strip_length = self.__pattern_min_row_length - self.__d_value + 1

        for i in range(0, len(self.__pattern_matrix)):
            for j in range(0, self.__strip_length):
                dgram_value = self.__pattern_matrix[i][j:j+self.__d_value]

                if i == len(self.__pattern_matrix) - 1:
                    self.__dgram_hor_pos_linked_list[j] = self.__dgram_hor_pos[dgram_value]
                    self.__dgram_hor_pos[dgram_value] = j
                elif self.__dgram_vertical_skip[dgram_value] > len(self.__pattern_matrix) - i - 1:
                    self.__dgram_vertical_skip[dgram_value] = len(self.__pattern_matrix) - i - 1

    def count_matches(self, text: list[str]):
        if not text:
            return 0

        text = text
        max_number_of_columns = max([len(row) for row in text])
        j = self.__strip_length - 1
        count = 0
        while j < max_number_of_columns - self.__pattern_min_row_length + self.__strip_length:
            i = len(self.__pattern_matrix) - 1

            while i < len(text):
                dgram_value = text[i][j:j + self.__d_value]
                k = self.__dgram_hor_pos[dgram_value]

                while k > -1:
                    count += self.__trivial_check_position(text, i - len(self.__pattern_matrix) + 1, j - k)
                    k = self.__dgram_hor_pos_linked_list[k]

                i = i + self.__dgram_vertical_skip[dgram_value]

            j = j + self.__strip_length

        return count

    def __trivial_check_position(self, text: list[str], i: int, j: int):
        if i < 0 or i >= len(text):
            return False
        if j < 0 or j >= len(text[i]):
            return False

        for pattern_i in range(0, len(self.__pattern_matrix)):
            for pattern_j in range(0, len(self.__pattern_matrix[pattern_i])):
                if (i + pattern_i >= len(text) or j + pattern_j >= len(text[i + pattern_i])
                        or text[i + pattern_i][j + pattern_j] != self.__pattern_matrix[pattern_i][pattern_j]):
                    return False

        return True

line0 = input().split(" ")
pw = int(line0[0])
ph = int(line0[1])
mh = int(line0[2])
mw = int(line0[3])

picture = []
masterpiece = []

for i in range(0, ph):
    picture.append(input())

for i in range(0, mh):
    masterpiece.append(input())

pattern_search = PatternSearch2D(picture)
pattern_count = pattern_search.count_matches(masterpiece)
print(pattern_count)
