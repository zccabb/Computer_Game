# -*- coding: utf-8 -*-
# author: MaoLongLong
# date: 2019/8/13
import numpy as np


class Move:
    def __init__(self, from_x, from_y, to_x, to_y):
        self.from_x = from_x
        self.from_y = from_y
        self.to_x = to_x
        self.to_y = to_y

    def __repr__(self):
        return 'from:[{}, {}]'.format(self.from_x, self.from_y) + \
               ' to:[{}, {}]'.format(self.to_x, self.to_y)


class State:
    red = 1
    blue = -1

    def __init__(self, board, next_to_move, key):
        self.board = board
        self.next_to_move = next_to_move
        self.key = key

    @property
    def game_result(self):
        if self.board[4][4] > 0 or np.sum(self.board < 0) == 0:
            return 1
        elif self.board[0][0] < 0 or np.sum(self.board > 0) == 0:
            return -1
        else:
            return None

    def is_game_over(self):
        return self.game_result is not None

    def move(self, move: Move):
        new_board = np.copy(self.board)
        new_board[move.to_x][move.to_y] = new_board[move.from_x][move.from_y]
        new_board[move.from_x][move.from_y] = 0
        next_to_move = State.red if self.next_to_move == State.blue \
            else State.blue
        return State(new_board, next_to_move, self.key)

    def _get_legal_actions(self, key):
        f = np.where(self.board == key * self.next_to_move)
        if len(f[0]) == 0:
            return None
        else:
            actions = []
            if self.next_to_move == State.red:
                if f[0][0] + 1 < 5:
                    actions.append(
                        Move(f[0][0], f[1][0], f[0][0] + 1, f[1][0]))
                if f[1][0] + 1 < 5:
                    actions.append(
                        Move(f[0][0], f[1][0], f[0][0], f[1][0] + 1))
                if f[0][0] + 1 < 5 and f[1][0] + 1 < 5:
                    actions.append(
                        Move(f[0][0], f[1][0], f[0][0] + 1, f[1][0] + 1))
            else:
                if f[0][0] - 1 >= 0:
                    actions.append(
                        Move(f[0][0], f[1][0], f[0][0] - 1, f[1][0]))
                if f[1][0] - 1 >= 0:
                    actions.append(
                        Move(f[0][0], f[1][0], f[0][0], f[1][0] - 1))
                if f[0][0] - 1 >= 0 and f[1][0] - 1 >= 0:
                    actions.append(
                        Move(f[0][0], f[1][0], f[0][0] - 1, f[1][0] - 1))
            return actions

    def get_legal_actions(self, key=None):
        key = self.key if key is None else key
        tmp = self._get_legal_actions(key)
        if tmp is not None:
            return tmp
        else:
            actions = []
            k = key - 1
            while k > 0:
                tmp = self._get_legal_actions(k)
                if tmp is not None:
                    actions += tmp
                    break
                k -= 1
            k = key + 1
            while k < 7:
                tmp = self._get_legal_actions(k)
                if tmp is not None:
                    actions += tmp
                    break
                k += 1
            return actions

# TEST CODE
# if __name__ == '__main__':
#     a = np.zeros((5, 5))
#     a[2][2] = 3
#     state = State(a, 6, 1)
#     print(state.get_legal_actions())
