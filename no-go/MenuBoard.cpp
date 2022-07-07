#include "MenuBoard.h"
char aa[9], bb[9];
int ff = 0;
MenuBoard::MenuBoard(Rect _pos)
{
	pos = _pos;
	img_board = 0;
	img_board = LoadTexture("img//menu//menuboard.bmp");
	game_state = MAIN_MENU;
	move_color = Color::BLACK;
}

void MenuBoard::Init()
{
	
}

void MenuBoard::Draw()
{
	DrawTexture(pos, img_board);
	Point tmp = Point(pos.lx, pos.ly) + Point(65, 100);
	DrawString("NoGo", "楷体", 58, tmp.x, tmp.y, 142.0 / 256, 70.0 / 256, 49.0 / 256);

	tmp = Point(pos.lx, pos.ly) + Point(20, 370);
	if (game_state == GAME_BLACK_WIN)
		DrawString("Black Win！", "楷体", 48, tmp.x, tmp.y, 0, 0, 0);
	if (game_state == GAME_WHITE_WIN)
		DrawString("White Win！", "楷体", 48, tmp.x, tmp.y, 1, 1, 1);

	tmp = Point(pos.lx, pos.ly) + Point(35, 310);
	if (game_state == IN_GAME)
	{
		if (ff == 2)
		{
			aa[0] = 'B';
			aa[1] = 'l';
			aa[2] = 'a';
			aa[3] = 'c';
			aa[4] = 'k';
			aa[5] = '(';
			bb[0] = 'W';
			bb[1] = 'h';
			bb[2] = 'i';
			bb[3] = 't';
			bb[4] = 'e';
			bb[5] = '(';
			if (move_color == Color::BLACK)
				DrawString(bb, "楷体", 40, tmp.x, tmp.y, 1, 1, 1);
			else
				DrawString(aa, "楷体", 40, tmp.x, tmp.y, 0, 0, 0);
		}
		else {
			if (move_color == Color::BLACK)
				DrawString(aa, "楷体", 40, tmp.x, tmp.y, 1, 1, 1);
			else
				DrawString(bb, "楷体", 40, tmp.x, tmp.y, 0, 0, 0);
		}
	}
}

void MenuBoard::OnMouseClick(Point a)
{
}

void MenuBoard::SetGameState(int id)
{
	game_state = id;
}

void MenuBoard::SetMoveColor(Color id)
{
	move_color = id;
}

bool MenuBoard::in(Point a)
{
	return pos.inRect(a);
}

void MenuBoard::xy(int x, int y)
{	
	aa[0] = 'W';
	aa[1] = 'h';
	aa[2] = 'i';
	aa[3] = 't';
	aa[4] = 'e';
	aa[5] = '(';
	char xx,yy;
	yy=((char)(y+ 'A'));
	xx = (char)((9 - x) + '0');
	aa[7] =xx;
	aa[6] = yy;
	aa[8] = ')';
	std::cout << aa[6]<<aa[7] << std::endl;
	
}

void MenuBoard::yx(int x, int y)
{
	
	bb[0] = 'B';
	bb[1] = 'l';
	bb[2] = 'a';
	bb[3] = 'c';
	bb[4] = 'k';
	bb[5] = '(';

	char xx, yy;
	xx = ((char)((9-x) + '0'));
	bb[7] = xx;
	yy = (char)(y + 'A' );
	bb[6] = yy;
	bb[8] = ')';
	std::cout << bb[6] << bb[7] << std::endl;
}
void MenuBoard::xxyy(int f)
{
	ff= f;
}