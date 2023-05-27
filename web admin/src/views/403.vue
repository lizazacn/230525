<template>
	<div class="error-page">
		<div class="error-code">C<span>h</span>e<span>r</span>r<span>y</span></div>
		<div class="error-desc">Cherry</div>

	</div>
</template>

<script setup lang="ts" name="403">
import { useRouter } from 'vue-router';

const router = useRouter();
const goBack = () => {
	router.go(-2);
};
</script>

<style scoped>
.error-page {
	display: flex;
	justify-content: center;
	align-items: center;
	flex-direction: column;
	width: 100%;
	height: 100%;
	background: #f3f3f3;
	box-sizing: border-box;
}
.error-code {
	line-height: 1;
	font-size: 250px;
	font-weight: bolder;
	color: #f02d2d;
}
.error-code span {
	color: #00a854;
}
.error-desc {
	font-size: 30px;
	color: #777;
}
.error-handle {
	margin-top: 30px;
	padding-bottom: 200px;
}
.error-btn {
	margin-left: 100px;
}
</style>
