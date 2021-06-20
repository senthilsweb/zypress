<!--
Reference : https://blog.devartis.com/custom-select-with-vue-js-39b010ddc1fb

-->
<template>
<div>
    <label :for="this.name" class="block text-sm font-medium text-gray-700">
        {{this.label}}
    </label>
    <div class="mt-1">
        <p v-if="$fetchState.pending">Fetching data...</p>
        <p v-else-if="$fetchState.error">An error occurred :(</p>
        <select @change="onChangeDropdown" v-else :id="this.name" :autocomplete="this.name" :name="this.name" class="h-full py-3 pl-4 pr-8 shadow-sm focus:ring-green-500 focus:border-green-500 block w-full text-sm font-medium border-gray-300 ">
            <option value="">Choose</option>
            <option v-for="(item, index) in this.items.data" v-bind:key="index" :value="item.code" :selected="selectedOption(item)">{{item.name}}</option>
        </select>
    </div>
</div>
</template>

<script>
export default {
    name: "dropdownlist",
    data() {
        return {
            items: this.data,
            selected_value: this.selected_value
        }
    },
    props: {
        type: String,
        label: String,
        name: String,
        selected_value: String,
        selected_text: String,
        api: String,
        data: Array
    },
    emits: ['selected_item'],
    methods: {
        selectedOption(option) {
            if (!s.isBlank(this.selected_value)) {
                return option.code.toString() === this.selected_value.toString();
            }
            return false;
        },
        onChangeDropdown(e) {
            const selectedCode = e.target.value;
            const option = this.items.data.find((option) => {
                return selectedCode === option.code;
            });
            try {
                this.selected_value = option.code;
                this.$emit("selected_item", option.code);
            } catch (e) {
                this.selected_value ="";
                this.$emit("selected_item", "");
            }

        },
        async getDropdown() {
            alert("getDropdown")
        },
        created() {

            //this.getDropdown()
        },
        beforeMount() {
            this.getDropdown()
        }
    },
    async fetch() {
        try {
            if (this.api !== undefined) {
                this.items = await fetch(this.$config.apiURL + this.api).then(res => res.json())
            } else if (this.data !== undefined) {
                this.items = this.data
            } else {
                this.items = ['No data']
            }
        } catch (error) {
            console.log(error)
        }
    }
}
</script>
