defmodule Day1 do
  def find_the_most_total_calories(logs) do
    logs
    |> String.split("\n\n")
    |> Enum.map(&String.trim/1)
    |> Enum.map(&sum_by_each_elf/1)
    |> Enum.reduce(0, &get_the_most_total_calories/2)
  end

  defp sum_by_each_elf(calories) do
    calories
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

input = Path.absname("day1.txt", "./input/2022") |> File.read!()


input
|> Day1.find_the_most_total_calories()
|> IO.puts()
