#include <iostream>

int main() {
  int N{}, M{};
  std::cin >> N >> M;

  int fail = 0;

  for (int i = 0; i < M; i++) {
    int group;
    std::cin >> group;

    int newN = N - group;
    if (newN < 0) {
      fail++;
      continue;
    }

    N = newN;
  }

  std::cout << fail << std::endl;
  return 0;
}
