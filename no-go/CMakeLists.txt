cmake_minimum_required(VERSION 3.19)
project(no_go)

set(CMAKE_CXX_STANDARD 14)

include_directories(.)
include_directories(Botzone_MCTS/jsoncpp)

add_executable(no_go
        Botzone_MCTS/jsoncpp/json-forwards.h
        Botzone_MCTS/jsoncpp/json.h
        Botzone_MCTS/jsoncpp.cpp
        Botzone_MCTS/main.cpp
        img/button/button1.bmp
        img/button/button2.bmp
        img/button/button3.bmp
        img/checkerboard/board0.bmp
        img/checkerboard/board1.bmp
        img/checkerboard/board2.bmp
        img/checkerboard/board3.bmp
        img/checkerboard/board4.bmp
        img/checkerboard/board5.bmp
        img/checkerboard/board6.bmp
        img/checkerboard/board7.bmp
        img/checkerboard/board8.bmp
        img/checkerboard/board9.bmp
        img/checkerboard/checkerboard.bmp
        img/menu/menuboard.bmp
        img/piece/blackpiece.bmp
        img/piece/whitepiece.bmp
        saves/1.save
        saves/2.save
        Button.cpp
        Button.h
        CheckerBoard.cpp
        CheckerBoard.h
        GameAI.cpp
        GameAI.h
        GameAI_MCTS.cpp
        GameAI_MCTS.h
        GameAI_Minmax.cpp
        GameAI_Minmax.h
        GameRule.cpp
        GameRule.h
        GameScene.cpp
        GameScene.h
        Graphic.cpp
        Graphic.h
        main.cpp
        MenuBoard.cpp
        MenuBoard.h
        MinMaxNode.cpp
        MinMaxNode.h
        safe_queue.h
        SaveManager.cpp
        SaveManager.h)
