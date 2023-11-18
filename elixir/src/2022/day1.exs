defmodule Day1 do
  def find_the_most_total_calories(each_sum) do
    each_sum
    |> Enum.reduce(0, &get_the_most_total_calories/2)
  end

  def get_sum_of_top_3_elves(each_sum) do
    each_sum
    |> Enum.sort(&(&1 >= &2))
    |> Enum.take(3)
    |> Enum.sum()
  end

  def sum_by_each_elf(logs) do
    logs
    |> String.split("\n\n")
    |> Enum.map(&String.trim/1)
    |> Enum.map(&calcurate_each_elf/1)
  end

  defp calcurate_each_elf(item) do
    item
    |> String.split("\n")
    |> Enum.map(fn n -> Integer.parse(n) |> elem(0) end)
    |> Enum.sum()
  end

  defp get_the_most_total_calories(calrories, acc) do
    if calrories > acc do
      calrories
    else
      acc
    end
  end
end

input = Path.absname("day1.txt", "./input/2022")
|> File.read!()
|> Day1.sum_by_each_elf()


input
|> Day1.find_the_most_total_calories()
|> IO.puts()

input
|> Day1.get_sum_of_top_3_elves()
|> IO.puts()
